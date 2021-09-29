package Adapter

import "log"

type AdvancedMediaPlayer interface {
	playVlc(fileName string)
	playMp4(fileName string)
}

type MediaPlayer interface {
	play(audioType string, fileName string)
}

type VlcPlayer struct {
}

func (vlcPlayer *VlcPlayer) playVlc(fileName string) {
	log.Println("Vlc is playing..." + fileName)
}

func (vlcPlayer *VlcPlayer) playMp4(fileName string) {

}

type Mp4Player struct {
}

func (mp4Player *Mp4Player) playVlc(fileName string) {

}

func (mp4Player *Mp4Player) playMp4(fileName string) {
	log.Println("Mp4 is playing..." + fileName)
}

type MediaAdapter struct {
	advancedMediaPlayer AdvancedMediaPlayer
}

func setMediaAdapter(audioType string) *MediaAdapter {
	if audioType == "vlc" {
		return &MediaAdapter{
			&VlcPlayer{},
		}
	} else if audioType == "mp4" {
		return &MediaAdapter{
			&Mp4Player{},
		}
	}
	return nil
}

func (mediaAdapter *MediaAdapter) play(audioType string, fileName string) {
	if audioType == "vlc" {
		mediaAdapter.advancedMediaPlayer.playVlc(fileName)
	} else if audioType == "mp4" {
		mediaAdapter.advancedMediaPlayer.playMp4(fileName)
	}
}

type AudioPlayer struct {
	mediaAdapter MediaAdapter
}

func (audioPlayer *AudioPlayer) play(audioType string, fileName string) {
	if audioType == "mp3" {
		log.Println("Playing mp3 file Name:" + fileName)
	} else if audioType == "vlc" || audioType == "mp4" {
		audioPlayer.mediaAdapter = *setMediaAdapter(audioType)
		audioPlayer.mediaAdapter.play(audioType, fileName)
	} else {
		log.Println("not supported")
	}
}
