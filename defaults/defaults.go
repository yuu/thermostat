package defaults

import (
        "os"
        "log"

        "github.com/pelletier/go-toml"
)

type Defaults interface {
        Load(data interface{})
        Save(data interface{})
}

type defaults struct {
        path string
}

func New(path string) Defaults {
        return &defaults{path}
}

func open(path string) *os.File {
        flag := os.O_RDWR | os.O_CREATE
        const perm os.FileMode = 0644
        file, err := os.OpenFile(path, flag, perm)
        if err != nil {
                log.Fatalf("open error: %v, file: %v", err, path)
                return nil
        }

        return file
}

func (d *defaults) encoder() (*os.File, *toml.Encoder) {
        file := open(d.path)
        if file == nil {
                return nil, nil
        }

        return file, toml.NewEncoder(file)
}

func (d *defaults) decoder() (*os.File, *toml.Decoder) {
        file := open(d.path)
        if file == nil {
                return nil, nil
        }

        return file, toml.NewDecoder(file)
}

func (d *defaults) Save(data interface{}) {
        file, encoder := d.encoder()
        if encoder == nil {
                return
        }
        defer file.Close()
        if err := encoder.Encode(data); err != nil {
                log.Fatalf("toml encode error: %v", err)
        }
}

func (d *defaults) Load(data interface{}) {
        file, decoder := d.decoder()
        if decoder == nil {
                return
        }
        defer file.Close()
        if err := decoder.Decode(data); err != nil {
                log.Fatalf("toml decode error: %v", err)
        }
}
