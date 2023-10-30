package command

import (
	"bytes"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestArgoCdCmd(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		envMap map[string]string
		want   string
	}{
		{
			envMap: map[string]string{},
		},
	}

	for _, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			for ek, ev := range testCase.envMap {
				t.Setenv(ek, ev)
			}
			cmd := rootCmd
			var b bytes.Buffer
			cmd.SetOut(&b)
			cmd.SetErr(&b)
			cmd.SetArgs([]string{"argocd"})
			cmd.Execute()

			assert.Equal(t, b.String(), testCase.want)
		})
	}
}
