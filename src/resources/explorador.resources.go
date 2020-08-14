package resources

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/dhowden/tag"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
)
type Metadata  struct {
Artist string `json:"artist"`
Album  string `json:"album"`
Title  string `json:"title"`
Track  string `json:"track"`
Year   string `json:"year"`
Genre string `json:"genre"`
}
type Explorer struct {
	Filename  string `json:"filename"`
	Hash      string `json:"hash"`
	Totalpath string `json:"totalpath"`
	Metadata  Metadata `json:"metadata"`
}


func List_Dir(path string) []Explorer {
	var lista []Explorer
	var files []string
	//root := "/home/pordefecto/Música/Rock"
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && filepath.Ext(path) == ".mp3" || filepath.Ext(path) == ".mp4" {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {

		filename := filepath.Base(file)
		hash, err := hash_file_sha1(file)
		if err != nil {
			panic(err)
		}
		totalpath := file

		m := Get_tags(file)
		explo := Explorer {
			Filename: filename,
			Hash: hash,
			Totalpath: totalpath,
			Metadata: m,
		}

		fmt.Println("print object")
		fmt.Println(explo)
		Get_tags(file)
		lista = append(lista,explo)
	}
	fmt.Println(lista)
	return lista
}

func hash_file_sha1(filePath string) (string, error) {
	//Initialize variable returnMD5String now in case an error has to be returned
	var returnSHA1String string

	//Open the filepath passed by the argument and check for any error
	file, err := os.Open(filePath)
	if err != nil {
		return returnSHA1String, err
	}

	//Tell the program to call the following function when the current function returns
	defer file.Close()

	//Open a new SHA1 hash interface to write to
	hash := sha1.New()

	//Copy the file in the hash interface and check for any error
	if _, err := io.Copy(hash, file); err != nil {
		return returnSHA1String, err
	}

	//Get the 20 bytes hash
	hashInBytes := hash.Sum(nil)[:20]

	//Convert the bytes to a string
	returnSHA1String = hex.EncodeToString(hashInBytes)

	return returnSHA1String, nil

}

func Get_tags(filepath string) Metadata{
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("error loading file: %v", err)
	}
	defer f.Close()
	m, err := tag.ReadFrom(f)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(m.Format()) // The detected format.

	log.Print(m.Picture())
	artist := m.Artist()
	album  := m.Album()
	title  := m.Title()
	track, _  := m.Track()
	year   := m.Year()
	genre  := m.Genre()
	yeardos := strconv.Itoa(year)
	trackdos := strconv.Itoa(track)
	var meta Metadata
	meta = Metadata{Artist: artist,Album: album,Title: title,Track:trackdos,Year:yeardos,Genre: genre}
	return meta
}

func Get_Cover(filepath string) ([]byte,string) {
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("error loading file: %v", err)
	}
	defer f.Close()
	m, err := tag.ReadFrom(f)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(m.Format()) // The detected format.
	log.Print(m.Picture())
	log.Print(m.Picture().MIMEType)
	log.Print(m.Picture().Type)
	log.Print(m.Picture().String())
	log.Print(m.Picture().Data)
	data := m.Picture().Data
	mime := m.Picture().MIMEType
	return data,mime
}
func Get_Track(filepath string) string {
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("error loading file: %v", err)
	}
	defer f.Close()
	m, err := tag.ReadFrom(f)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(m.Lyrics())
	lyrics := m.Lyrics()
	return lyrics
}

