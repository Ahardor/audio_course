import pyaudio
import wave
import time
import keyboard
# import librosa
import numpy as np
import scipy
# from sound_to_midi.monophonic import wave_to_midi

CHUNK = 1024
FORMAT = pyaudio.paFloat32
CHANNELS = 2
RATE = 48000
RECORD_SECONDS = 5
WAVE_OUTPUT_FILENAME = "output.wav"

p = pyaudio.PyAudio()
stream: pyaudio.Stream = None
frames = []

def main():
    global stream
    global frames
    global p

    info = p.get_host_api_info_by_index(0)
    numdevices = info.get('deviceCount')

    for i in range(0, numdevices):
        if (p.get_device_info_by_host_api_device_index(0, i).get('maxInputChannels')) > 0:
            print("Input Device id ", i, " - ", p.get_device_info_by_host_api_device_index(0, i).get('name'))

    input_device = int(input("Enter input device id: "))

    stream = p.open(format=FORMAT,
                    channels=1,
                    rate=RATE,
                    input=True,
                    frames_per_buffer=CHUNK,
                    input_device_index=input_device)

    for i in range(5):
        print("recording in", 5 - i, "seconds")
        time.sleep(1)

    print("* recording")

    while True:
        data = stream.read(CHUNK)
        frames.append(data)

        peakfreq=0.0
        if(stream):
            tdarray=np.frombuffer(data, dtype=np.float32)

            fftarray=np.fft.fft(tdarray)

            nyquist=.5*RATE

            denom,numer=scipy.signal.butter(5,[4/nyquist, 10000/nyquist],btype="band")  
            ffftarray=scipy.signal.lfilter(denom,numer,fftarray)

            freqarray=np.fft.fftfreq(tdarray.size)

            ind=np.argmax(np.abs(ffftarray)**2)

            peakfreq=np.abs(freqarray[ind])*RATE
            
            frequency=peakfreq

            print(frequency)

        if keyboard.is_pressed('q'):
            print("* done recording")

            stream.stop_stream()
            stream.close()
            p.terminate()

            wf = wave.open(WAVE_OUTPUT_FILENAME, 'wb')
            wf.setnchannels(CHANNELS)
            wf.setsampwidth(p.get_sample_size(FORMAT))
            wf.setframerate(RATE)
            wf.writeframes(b''.join(frames))
            wf.close()

            
            # x = np.frombuffer(b''.join(frames), dtype=np.int16)
            # test = librosa.feature.mfcc(y=x.astype(np.float16), sr=RATE, )
            # midi = wave_to_midi(audio_signal=test.astype(np.float16), srate=RATE, frame_length=CHUNK)
        
            # print(test.astype(np.float32))
            
            # with open ("output.mid", 'wb') as f:
            #     midi.writeFile(f)

            break

main()