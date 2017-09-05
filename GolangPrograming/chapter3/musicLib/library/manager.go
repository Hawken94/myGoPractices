package library

import (
	"errors"
)

// Music 音乐结构体
type Music struct {
	Id     string
	Name   string
	Artist string
	Source string
	Type   string
}

// MusicManager 音乐管理
type MusicManager struct {
	musics []Music
}

// NewMusicManager 初始化
func NewMusicManager() *MusicManager {
	return &MusicManager{make([]Music, 0)}
}

func (m *MusicManager) Len() int {
	return len(m.musics)
}
func (m *MusicManager) Get(index int) (*Music, error) {
	if index < 0 || index >= len(m.musics) {
		return nil, errors.New("Index out of range.")
	}
	return &m.musics[index], nil
}

func (m *MusicManager) Find(name string) *Music {
	if len(m.musics) == 0 {
		return nil
	}
	for _, m := range m.musics {
		if m.Name == name {
			return &m
		}
	}
	return nil
}

func (m *MusicManager) Add(music *Music) {
	m.musics = append(m.musics, *music)
}

func (m *MusicManager) Remove(index int) *Music {
	if index < 0 || index >= len(m.musics) {
		return nil
	}

	removeMusic := &m.musics[index]
	// 从数组切片中删除元素
	if index < len(m.musics)-1 { // 中间元素
		m.musics = append(m.musics[:index-1], m.musics[index+1:]...)
	} else if index == 0 { // 删除仅有的一个元素 todo ???
		m.musics = make([]Music, 0)
	} else { // 删除的是最后一个元素
		m.musics = m.musics[:index-1]
	}

	return removeMusic
}

func (m *MusicManager) RemoveByName(name string) *Music {
	if name == "" {
		return nil
	}
	var removeMusic *Music
	// 只寻找一次
	for i := 0; i < len(m.musics); i++ {
		if m.musics[i].Name == name {
			removeMusic = &m.musics[i]
			if len(m.musics) > 1 { // todo 这样写有很大的问题
				m.musics = append(m.musics[:i-1], m.musics[i+1:]...)
			}
			break
		}
	}

	return removeMusic
}
