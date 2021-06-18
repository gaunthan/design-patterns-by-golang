package bridge

import "testing"

func TestSendMessage(t *testing.T) {
	cases := []struct {
		name    string
		mail    Mail
		mailman Mailman
		want    string
	}{
		{
			name: "common_mailman_dog",
			mail: Mail{
				from:    "Tom",
				to:      "Jerry",
				content: "Hi",
			},
			mailman: NewCommonMailman(NewDogSender()),
			want:    "[Dog] [ ] Tom -> Jerry: Hi",
		},
		{
			name: "common_mailman_cat",
			mail: Mail{
				from:    "Tom",
				to:      "Jerry",
				content: "Hi",
			},
			mailman: NewCommonMailman(NewCatSender()),
			want:    "[Cat] [ ] Tom -> Jerry: Hi",
		},
		{
			name: "special_mailman_dog",
			mail: Mail{
				from:    "Tom",
				to:      "Jerry",
				content: "Hi",
			},
			mailman: NewSpecialMailman(NewDogSender()),
			want:    "[Dog] [*] Tom -> Jerry: Hi",
		},
		{
			name: "special_mailman_cat",
			mail: Mail{
				from:    "Tom",
				to:      "Jerry",
				content: "Hi",
			},
			mailman: NewSpecialMailman(NewCatSender()),
			want:    "[Cat] [*] Tom -> Jerry: Hi",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.mailman.SendMail(tt.mail)
			if got != tt.want {
				t.Errorf("got %q want %q", got, tt.want)
			}
		})
	}
}
