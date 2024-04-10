package extractembed

import (
	"embed"
	"os"
)

func ExtractTemp(embedFs *embed.FS) (string, error) {
	dname, err := os.MkdirTemp("", "sampledir")
	if err != nil {
		return "", err
	}

	return dname, Extract(dname, embedFs)
}

func Extract(base string, embedFs *embed.FS) error {
	embedFilePaths, err := listEmbedFiles(".", embedFs)
	if err != nil {
		return err
	}

	for _, ef := range embedFilePaths {
		body, err := readEmbedFile(ef, embedFs)
		if err != nil {
			return err
		}

		makeSuperFolder(ef, base)

		err = saveFile(ef, body, base)
		if err != nil {
			return err
		}
	}

	return nil
}
