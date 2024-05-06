import copy
import os
import numpy as np
import scipy.fftpack
import sounddevice as sd
import time

import eel

deviceIndex = -1
# General settings
SAMPLE_FREQ = 48000 # sample frequency in Hz
WINDOW_SIZE = 4800 # window size of the DFT in samples
WINDOW_STEP = 2400 # step size of window
NUM_HPS = 5 # max number of harmonic product spectrums
POWER_THRESH = 4e-4 # tuning is activated if the signal power exceeds this threshold
CONCERT_PITCH = 440 # defining a1
WHITE_NOISE_THRESH = 0.2 # everything under WHITE_NOISE_THRESH*avg_energy_per_freq is cut off

WINDOW_T_LEN = WINDOW_SIZE / SAMPLE_FREQ # length of the window in seconds
SAMPLE_T_LENGTH = 1 / SAMPLE_FREQ # length between two samples in seconds
DELTA_FREQ = SAMPLE_FREQ / WINDOW_SIZE # frequency step width of the interpolated DFT
OCTAVE_BANDS = [50, 100, 200, 400, 800, 1600, 3200, 6400, 12800, 25600]
# This function finds the closest note for a given pitch
# Returns: note (e.g. A4, G#3, ..), pitch of the tone
CONCERT_PITCH = 440
ALL_NOTES = ["A","A#","B","C","C#","D","D#","E","F","F#","G","G#"]
def find_closest_note(pitch):
  i = int(np.round(np.log2(pitch/CONCERT_PITCH)*12))
  closest_note = ALL_NOTES[i%12] + str(4 + (i + 9) // 12)
  closest_pitch = CONCERT_PITCH*2**(i/12)
  return closest_note, closest_pitch


HANN_WINDOW = np.hanning(WINDOW_SIZE)

old_closest_note = 'A0'


# The sounddecive callback function
# Provides us with new data once WINDOW_STEP samples have been fetched
def callback(indata, frames, time, status):
  global old_closest_note
  if not hasattr(callback, "window_samples"):
    callback.window_samples = [0 for _ in range(WINDOW_SIZE)]
  if not hasattr(callback, "noteBuffer"):
    callback.noteBuffer = ["1","2"]

  if status:
    print(status)
  if any(indata):
    callback.window_samples = np.concatenate((callback.window_samples, indata[:, 0])) # append new samples
    callback.window_samples = callback.window_samples[len(indata[:, 0]):] # remove old samples

    signal_power = (np.linalg.norm(callback.window_samples, ord=2)**2) / len(callback.window_samples)
    if signal_power < POWER_THRESH:
      # os.system('cls' if os.name=='nt' else 'clear')
      if(old_closest_note != '...'):
        # print("Closest note: ...")
        old_closest_note = '...'
      return
    
    hann_samples = callback.window_samples * HANN_WINDOW
    magnitude_spec = abs(scipy.fftpack.fft(hann_samples)[:len(hann_samples)//2])

    for i in range(int(62/(SAMPLE_FREQ/WINDOW_SIZE))):
      magnitude_spec[i] = 0 #suppress mains hum

    # calculate average energy per frequency for the octave bands
    # and suppress everything below it
    for j in range(len(OCTAVE_BANDS)-1):
      ind_start = int(OCTAVE_BANDS[j]/DELTA_FREQ)
      ind_end = int(OCTAVE_BANDS[j+1]/DELTA_FREQ)
      ind_end = ind_end if len(magnitude_spec) > ind_end else len(magnitude_spec)
      avg_energy_per_freq = (np.linalg.norm(magnitude_spec[ind_start:ind_end], ord=2)**2) / (ind_end-ind_start)
      avg_energy_per_freq = avg_energy_per_freq**0.5
      for i in range(ind_start, ind_end):
        magnitude_spec[i] = magnitude_spec[i] if magnitude_spec[i] > WHITE_NOISE_THRESH*avg_energy_per_freq else 0

    # interpolate spectrum
    mag_spec_ipol = np.interp(np.arange(0, len(magnitude_spec), 1/NUM_HPS), np.arange(0, len(magnitude_spec)),
                              magnitude_spec)
    mag_spec_ipol = mag_spec_ipol / np.linalg.norm(mag_spec_ipol, ord=2) #normalize it

    hps_spec = copy.deepcopy(mag_spec_ipol)

    # calculate the HPS
    for i in range(NUM_HPS):
      tmp_hps_spec = np.multiply(hps_spec[:int(np.ceil(len(mag_spec_ipol)/(i+1)))], mag_spec_ipol[::(i+1)])
      if not any(tmp_hps_spec):
        break
      hps_spec = tmp_hps_spec

    max_ind = np.argmax(hps_spec)
    max_freq = max_ind * (SAMPLE_FREQ/WINDOW_SIZE) / NUM_HPS

    closest_note, closest_pitch = find_closest_note(max_freq)
    max_freq = round(max_freq, 1)
    closest_pitch = round(closest_pitch, 1)

    callback.noteBuffer.insert(0, closest_note) # note that this is a ringbuffer
    callback.noteBuffer.pop()

    # os.system('cls' if os.name=='nt' else 'clear')
    # if callback.noteBuffer.count(callback.noteBuffer[0]) == len(callback.noteBuffer):
    if old_closest_note != closest_note:
      old_closest_note = closest_note
      eel.changeNote(closest_note)
      print(f"Closest note: {closest_note} {max_freq}/{closest_pitch}")
    # else:
    #   print(f"Closest note: ...")
  else:
    print('no input')

while deviceIndex == -1:
  arr = []
  print("Select device:")
  for i in sd.query_devices():
    if i['max_input_channels'] > 0 and i['hostapi']==3:
      print(i['index'], i['name'])
      arr.append(i['index'])

  ind = int(input())
  if ind in arr:
    deviceIndex = ind
    break

def detect():
  try:
    with sd.InputStream(channels=1, callback=callback,
      blocksize=WINDOW_STEP,
      samplerate=SAMPLE_FREQ, device=deviceIndex):
      while True:
        eel.sleep(1.0)
        pass
  except Exception as e:
    print(str(e))

eel.init('ClientPart/front')
eel.spawn(detect)
eel.start('index.html', block=False)

# Start the microphone input stream
while True:
    # print("I'm a main loop")
    eel.sleep(1.0)    

