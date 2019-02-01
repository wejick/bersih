package service

import (
	"context"

	"github.com/wejick/bersih/pkg/autocomplete/model"
)

// TextDataList list of text
type TextDataList struct {
	TotalData int64
	Text      []TextData
}

// TextData data block of text
type TextData struct {
	Data string `json:"data"`
	URL  string `json:"url"`
}

// GetText returns list of text
func (S *Service) GetText(ctx context.Context) (textDataList TextDataList, err error) {
	textList, err := S.repo.GetText(ctx)
	if err != nil {
		return
	}
	for _, text := range textList.Data {
		textDataList.Text = append(textDataList.Text, textToTextData(text))
	}
	textDataList.TotalData = textList.TotalData

	return
}

func textDataToModel(textData TextData) (text model.Text) {
	text = model.Text{
		Data: textData.Data,
		URL:  textData.URL,
	}
	return
}

func textToTextData(text model.Text) (textData TextData) {
	textData = TextData{
		Data: text.Data,
		URL:  text.URL,
	}
	return
}
