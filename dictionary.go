package yandextranslate

import (
	dictionary "github.com/dafanasev/go-yandex-dictionary"
)

type Translate struct {
	Eng           string
	Rus           string
	Transcription string
}

func TranslateRus(rus, api string) (*Translate, error) {
	dict := dictionary.New(api)
	_, err := dict.GetLangs()
	if err != nil {
		return nil, err
	}
	definition, err := dict.Lookup(&dictionary.Params{Lang: "ru-en", Text: rus})

	if err != nil {
		return nil, err
	}

	ret, err := TranslateEng(definition.Def[0].Tr[0].Text, api)

	return ret, err
}

func TranslateEng(eng, api string) (*Translate, error) {
	dict := dictionary.New(api)
	_, err := dict.GetLangs()
	if err != nil {
		return nil, err
	}
	definition, err := dict.Lookup(&dictionary.Params{Lang: "en-ru", Text: eng})
	if err != nil {
		return nil, err
	}
	var ret = Translate{
		Eng:           definition.Def[0].Tr[0].Text,
		Rus:           eng,
		Transcription: definition.Def[0].Ts,
	}

	return &ret, err
}
