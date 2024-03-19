import pyaudio
import wave
import time
import keyboard
import numpy as np
import matplotlib.pyplot as plt

CHUNK = 1024
FORMAT = pyaudio.paInt16
CHANNELS = 2
RATE = 44100
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
                    channels=CHANNELS,
                    rate=RATE,
                    input=True,
                    frames_per_buffer=CHUNK,
                    input_device_index=input_device)

    for i in range(5):
        print("recording in", 5 - i, "seconds")
        time.sleep(1)

    print("* recording")

    result = np.array([], dtype=np.int16)
    
    while True:
        data = stream.read(CHUNK)
        frames.append(data)
        newdata = np.fromstring(data, dtype=np.int16)
        result = np.append(result, newdata)
        

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

            

            break

main()