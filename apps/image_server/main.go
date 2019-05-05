package main

import (
	"github.com/arunscape/friends/commons/server/logger"
	"github.com/arunscape/friends/commons/server/utils"

	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var (
	STORAGE = os.Getenv("IMAGE_DIRECTORY")
)

const VIEW_ENDPOINT = "/view/"

func main() {
	if os.Getenv("DID_I_SET_THE_ENVIROMENT_VARIABLES") != "YES I DID" {
		logger.Error("Enviroment variables not found")
		return
	}
	os.Mkdir(STORAGE, os.ModeDir|0777)
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc(VIEW_ENDPOINT, viewHandler)

	port := 8080
	port_env, _ := strconv.Atoi(os.Getenv("PORT"))
	if port_env != 0 {
		port = port_env
	}
	logger.Info("Starting Server (PORT " + strconv.Itoa(port) + ")")
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	logger.Debug(r)
	var err error

	/*
		    preImage, err := r.GetBody()
		    if err != nil {
				logger.Error(err)
				w.WriteHeader(400)
				w.Write([]byte("400 - Bad Request"))
		        return
		    }

		    err = preprocessImage(preImage)
			if err != nil {
				logger.Error(err)
				w.WriteHeader(400)
				w.Write([]byte("400 - Bad Request"))
				return
			}
	*/

	path := generatePath()
	err = setImage(path, r.Body)
	if err != nil {
		logger.Error(err)
		w.WriteHeader(400)
		w.Write([]byte("400 - Bad Request"))
		return
	}

	w.Write([]byte(path))
	logger.Info("Uploaded Image: " + path)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	logger.Debug(r)
	image_path := STORAGE + strings.TrimPrefix(r.URL.String(), VIEW_ENDPOINT)
	logger.Info("Retriving image: ", image_path)

	err := getImage(image_path, w)

	if err != nil {
		logger.Error("Could not retrieve image at path " + image_path)
		w.WriteHeader(404)
	}
}

func preprocessImage(img io.ReadCloser) error {
	// TODO: downscale image maybe? Verify it indeed is an image file
	return nil
}

func generatePath() string {
	return STORAGE + utils.UUID()
}

func setImage(path string, image io.ReadCloser) error {
	logger.Debug("Saving to path ", path)
	file, err := os.Create(path)
	if err != nil {
		logger.Error("Could not Create file at path: ", path)
		return err
	}
	return copyData(image, file)
}

func getImage(path string, resp io.Writer) error {
	file, err := os.Open(path)
	if err != nil {
		logger.Error("Could not open file at path: ", path)
		return err
	}
	return copyData(file, resp)
}

func copyData(r io.Reader, w io.Writer) error {
	var err error
	buf := make([]byte, 20480)
	size := 0
	bytes := 0
	for err == nil {
		bytes, err = r.Read(buf)
		w.Write(buf[:bytes])
		size += bytes
		logger.Debug("Wrote ", size, " bytes so far")
	}
	if err != io.EOF {
		return err
	}
	return nil
}
