package utils

import (
    "context"
    "github.com/cloudinary/cloudinary-go/v2"
    "github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadImage(filePath string) (string, error) {
    cld, _ := cloudinary.NewFromParams("dquhhywhh", "174298959362779", "nmrHi3PugXbjBUuNqTmYhjTKQ2s")

    resp, err := cld.Upload.Upload(context.Background(), filePath, uploader.UploadParams{})
    if err != nil {
        return "", err
    }

    return resp.SecureURL, nil
} 