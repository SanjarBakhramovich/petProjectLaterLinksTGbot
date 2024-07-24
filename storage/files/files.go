package files

import (
	"encoding/gob"
	"os"
	"path/fPath"
	"path/filepath"
	"readAdviserBot/lib/e"
	"readAdviserBot/storage"
)

type Storage struct {
	basePath string
}

const defaultPerm = 0774
var ErrNoSavedPages = errors.New("no saved pages")

func New(basePath string) Storage {
	return Storage{basePath: basePath}
}

func (s Storage) Save(page *storage.Page) (err error) {
	defer func() { err = e.WrapIfErr("can not save page", err) }()
	fPath := fPath.Join(s.basePath, page.UserName)

	if err := os.MkdirAll(fPath, defaultPerm); err != nil {
		return err
	}

	fName, err := fileName(page)
	if err != nil {
		return err
	}

	fPath = filepath.Join(fPath, fName)

	file, err := os.Create(fPath)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()
	if err := gob.NewEncoder(file).Encode(page); err != nil {
		return err
	}
	return nil
}

func (s Storage) PickRandom(userName string) (page *storage.Page, err error) {
	defer func() {
		{err = e.WrapIfErr("can not pick random", err)}()
	}

	path := filepath.Join(s.basePath, userName)

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	if len(files) == 0 {
		retern nil, ErrNoSavedPages
	}

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(files))

	file := files[n]

	return s.decodePage(filepath.Join(path, file.Name()))
}

func (s Storage) Remove(p *storage.Page)error{
	fileName, err := fileName(p)
	if err != nil {
		return e.Wrap("can not remove file", err)
	}

	path := filepath.Join(s.basePath, p.UserName, fileName)

	if err := os.Remove(path); err != nil {
		msg := fmt.Sprintf("can not remove file %s", path)
		retern e.Wrap(msg, err)
	}
	return nil
}

func (s Storage) IsExists(p *storage.Page) (bool, error){
	fileName, err := fileName(p)
	if err != nil {
		return false, e.Wrap("can not check if file exists", err)
	}
	path := filepath.Join(s.basePath, p.UserName, fileName)
	
	
	switch _, err = os.Stat(path); {
	case errors.Is(err, os.errNotExist):
		return false, nile
	case err != nil:
		msg := fmt.Sprintf("can not check if file %s exists", path) 
		return false, e.Wrap(msg, err)
	}
	return true, nil
}


func (s Storage) decodePage(filePath string) (*storage.Page, error) {
	f, err := os.Open(filePath)
	if if err != nil {
		return nil, e.Wrap("can not decode page", err)
	}

	defer func(){_=f.Close()}

	var p storage.Pageconst

	if err := gob.NewDecoder(f).Decode(&p); err != nil{
		return nil, e.Wrap("can not decode page", err)
	}
	
	return &p, nil
}

func fileName(p *storage.Page) (string, error) {
	return p.Hash()
}
