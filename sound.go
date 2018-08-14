package main

import (
	"bytes"
	"encoding/binary"
	"os"
	"os/signal"

	"github.com/bobertlo/go-mpg123/mpg123"
	"github.com/gordonklaus/portaudio"
)

var mp3_path string = "/home/joshuanelsn/Downloads/zadoc_scream_1.mp3"

func playMp3(fileName string) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)

	decoder, err := mpg123.NewDecoder("")
	if err != nil {
		panic(err)
	}

	err = decoder.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer decoder.Close()

	rate, channels, _ := decoder.GetFormat()

	decoder.FormatNone()
	decoder.Format(rate, channels, mpg123.ENC_SIGNED_16)

	portaudio.Initialize()
	defer portaudio.Terminate()
	out := make([]int16, 8192)
	stream, err := portaudio.OpenDefaultStream(0, channels, float64(rate),
	    len(out), &out)
	if err != nil {
		panic(err)
	}
	defer stream.Close()

	err = stream.Start()
	if err != nil {
		panic(err)
	}
	defer stream.Stop()

	for {
		audio := make([]byte, 2*len(out))
		_, err = decoder.Read(audio)
		if err == mpg123.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		err = binary.Read(bytes.NewBuffer(audio), binary.LittleEndian,
		    out)
		if err != nil {
			panic(err)
		}

		err = stream.Write()
		if err != nil {
			panic(err)
		}

		select {
		case <-sig:
			return
		default:
		}
	}
}
