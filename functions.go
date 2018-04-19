package voiceit2

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

// Alternative initialization method which returns the initialized object rather than operating on an instance of struct.
func NewVoiceIt2(api_key, api_token string) VoiceIt2 {
	return VoiceIt2{
		APIKey:   api_key,
		APIToken: api_token,
		BaseURL:  "https://api.voiceit.io",
	}
}

// Initialization
// First assign the API Credentials.
func (vi *VoiceIt2) SetAPIKeyAndToken(key string, tok string) {
	vi.APIKey = key
	vi.APIToken = tok
}

// Get all users associated with the apiKey
func (vi *VoiceIt2) GetAllUsers() *bytes.Buffer {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.Close()

	req, _ := http.NewRequest("GET", vi.BaseURL+"/users", body)
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	respBody := &bytes.Buffer{}
	respBody.ReadFrom(resp.Body)

	return respBody
}

//Create a new user
func (vi *VoiceIt2) CreateUser() *bytes.Buffer {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseURL+"/users", body)
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	respBody := &bytes.Buffer{}
	respBody.ReadFrom(resp.Body)

	return respBody
}

// Check whether a user exists for the given
//	userId
// (begins with 'usr_')
func (vi *VoiceIt2) GetUser(userId string) *bytes.Buffer {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.Close()

	req, _ := http.NewRequest("GET", vi.BaseURL+"/users/"+userId, body)
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	respBody := &bytes.Buffer{}
	respBody.ReadFrom(resp.Body)

	return respBody
}

// Delete user with given
//	userId
// (begins with 'usr_')
func (vi *VoiceIt2) DeleteUser(userId string) *bytes.Buffer {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.Close()

	req, _ := http.NewRequest("DELETE", vi.BaseURL+"/users/"+userId, body)
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	respBody := &bytes.Buffer{}
	respBody.ReadFrom(resp.Body)

	return respBody
}

// Get a list of groups that the user with given
//	userId
// (begins with 'usr_') is a part of
func (vi *VoiceIt2) GetGroupsForUser(userId string) *bytes.Buffer {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.Close()

	req, _ := http.NewRequest("GET", vi.BaseURL+"/users/"+userId+"/groups", body)
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	respBody := &bytes.Buffer{}
	respBody.ReadFrom(resp.Body)

	return respBody
}

// Gets all enrollment for user with given
//	userId
// (begins with 'usr_')
func (vi *VoiceIt2) GetAllEnrollmentsForUser(userId string) *bytes.Buffer {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.Close()

	req, _ := http.NewRequest("GET", vi.BaseURL+"/enrollments/"+userId, body)
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	respBody := &bytes.Buffer{}
	respBody.ReadFrom(resp.Body)

	return respBody
}

func (vi *VoiceIt2) DeleteAllEnrollmentsForUser(userId string) *bytes.Buffer {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.Close()

	req, _ := http.NewRequest("DELETE", vi.BaseURL+"/enrollments/"+userId+"/all", body)
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	respBody := &bytes.Buffer{}
	respBody.ReadFrom(resp.Body)

	return respBody
}

func (vi *VoiceIt2) GetFaceFaceEnrollmentsForUser(userId string) *bytes.Buffer {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.Close()

	req, _ := http.NewRequest("GET", vi.BaseURL+"/enrollments/face/"+userId, body)
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	respBody := &bytes.Buffer{}
	respBody.ReadFrom(resp.Body)

	return respBody
}

// Create audio enrollment for user with given
//	userId(begins with 'usr_')
// and
//	contentLanguage
// ('en-US','es-ES', etc.)
// Note: File recording need to be no less than 1.2 seconds and no more than 5 seconds
func (vi *VoiceIt2) CreateVoiceEnrollment(userId string, contentLanguage string, filePath string) *bytes.Buffer {

	file, _ := os.Open(filePath)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, _ := writer.CreateFormFile("recording", filepath.Base(file.Name()))
	io.Copy(part, file)
	_ = writer.WriteField("userId", userId)
	_ = writer.WriteField("contentLanguage", contentLanguage)
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseURL+"/enrollments", body)
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	respBody := &bytes.Buffer{}
	respBody.ReadFrom(resp.Body)

	return respBody
}

func (vi *VoiceIt2) CreateVoiceEnrollmentByUrl(userId string, contentLanguage string, fileUrl string) *bytes.Buffer {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	_ = writer.WriteField("userId", userId)
	_ = writer.WriteField("contentLanguage", contentLanguage)
	_ = writer.WriteField("fileUrl", fileUrl)
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseURL+"/enrollments/byUrl", body)
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	respBody := &bytes.Buffer{}
	respBody.ReadFrom(resp.Body)

	return respBody
}

func (vi *VoiceIt2) CreateFaceEnrollment(userId string, contentLanguage string, filePath string, doBlinkDetection ...bool) *bytes.Buffer {

	file, _ := os.Open(filePath)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, _ := writer.CreateFormFile("video", filepath.Base(file.Name()))
	io.Copy(part, file)
	_ = writer.WriteField("userId", userId)
	_ = writer.WriteField("contentLanguage", contentLanguage)
	if len(doBlinkDetection) > 0 {
		_ = writer.WriteField("doBlinkDetection", strconv.FormatBool(doBlinkDetection[0]))
	}
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseURL+"/enrollments/face", body)
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	respBody := &bytes.Buffer{}
	respBody.ReadFrom(resp.Body)

	return respBody
}

// Create video enrollment for user with given
//	userId
// (begins with 'usr_')
// and
//	contentLanguage
// ('en-US','es-ES', etc.)
// Note: File recording need to be no less than 1.2 seconds and no more than 5 seconds
func (vi *VoiceIt2) CreateVideoEnrollment(userId string, contentLanguage string, filePath string, doBlinkDetection ...bool) *bytes.Buffer {

	file, _ := os.Open(filePath)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, _ := writer.CreateFormFile("video", filepath.Base(file.Name()))
	io.Copy(part, file)
	_ = writer.WriteField("userId", userId)
	_ = writer.WriteField("contentLanguage", contentLanguage)
	if len(doBlinkDetection) > 0 {
		_ = writer.WriteField("doBlinkDetection", strconv.FormatBool(doBlinkDetection[0]))
	}
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseURL+"/enrollments/video", body)
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	respBody := &bytes.Buffer{}
	respBody.ReadFrom(resp.Body)

	return respBody
}

func (vi *VoiceIt2) CreateVideoEnrollmentByUrl(userId string, contentLanguage string, fileUrl string, doBlinkDetection ...bool) *bytes.Buffer {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	_ = writer.WriteField("userId", userId)
	_ = writer.WriteField("contentLanguage", contentLanguage)
	_ = writer.WriteField("fileUrl", fileUrl)
	if len(doBlinkDetection) > 0 {
		_ = writer.WriteField("doBlinkDetection", strconv.FormatBool(doBlinkDetection[0]))
	}
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseURL+"/enrollments/video/byUrl", body)
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	respBody := &bytes.Buffer{}
	respBody.ReadFrom(resp.Body)

	return respBody
}

func (vi *VoiceIt2) DeleteFaceEnrollment(userId string, faceEnrollmentId string) *bytes.Buffer {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.Close()

	req, _ := http.NewRequest("DELETE", vi.BaseURL+"/enrollments/face/"+userId+"/"+faceEnrollmentId, body)
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	respBody := &bytes.Buffer{}
	respBody.ReadFrom(resp.Body)

	return respBody
}

// Delete enrollment for user with given
//	userId
// (begins with 'usr_') and
//	enrollmentId
// (integer)
func (vi *VoiceIt2) DeleteEnrollmentForUser(userId string, faceEnrollmentId string) *bytes.Buffer {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.Close()

	req, _ := http.NewRequest("DELETE", vi.BaseURL+"/enrollments/"+userId+"/"+faceEnrollmentId, body)
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	respBody := &bytes.Buffer{}
	respBody.ReadFrom(resp.Body)

	return respBody
}

// Get all the groups associated with the apiKey
func (vi *VoiceIt2) GetAllGroups() *bytes.Buffer {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.Close()

	req, _ := http.NewRequest("GET", vi.BaseURL+"/groups", body)
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	respBody := &bytes.Buffer{}
	respBody.ReadFrom(resp.Body)

	return respBody
}

// Returns a group for the given
//	groupId
// (begins with 'grp_')
func (vi *VoiceIt2) GetGroup(groupId string) *bytes.Buffer {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.Close()

	req, _ := http.NewRequest("GET", vi.BaseURL+"/groups/"+groupId, body)
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	respBody := &bytes.Buffer{}
	respBody.ReadFrom(resp.Body)

	return respBody
}

// Checks if group with given
//	groupId
// (begins with 'grp_') exists
func (vi *VoiceIt2) GroupExists(groupId string) *bytes.Buffer {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.Close()

	req, _ := http.NewRequest("GET", vi.BaseURL+"/groups/"+groupId+"/exists", body)
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	respBody := &bytes.Buffer{}
	respBody.ReadFrom(resp.Body)

	return respBody
}

// Create a new group with the given description
func (vi *VoiceIt2) CreateGroup(description string) *bytes.Buffer {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	_ = writer.WriteField("description", description)

	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseURL+"/groups", body)
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	respBody := &bytes.Buffer{}
	respBody.ReadFrom(resp.Body)

	return respBody
}

// Adds user with given
//	userId
// (begins with 'usr_') to group with given
//	groupId
// (begins with 'grp_')
func (vi *VoiceIt2) AddUserToGroup(groupId string, userId string) *bytes.Buffer {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	_ = writer.WriteField("groupId", groupId)
	_ = writer.WriteField("userId", userId)

	writer.Close()

	req, _ := http.NewRequest("PUT", vi.BaseURL+"/groups/addUser", body)
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	respBody := &bytes.Buffer{}
	respBody.ReadFrom(resp.Body)

	return respBody
}

// Removes user with given
//	userId
// (begins with 'usr_') from group with given
//	groupId
// (begins with 'grp_')
func (vi *VoiceIt2) RemoveUserFromGroup(groupId string, userId string) *bytes.Buffer {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	_ = writer.WriteField("groupId", groupId)
	_ = writer.WriteField("userId", userId)

	writer.Close()

	req, _ := http.NewRequest("PUT", vi.BaseURL+"/groups/removeUser", body)
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	respBody := &bytes.Buffer{}
	respBody.ReadFrom(resp.Body)

	return respBody
}

// Delete group with given
//	groupId
// (begins with 'grp_')
// Note: This call does not delete any users, but simply deletes the group and disassociates the users from the group
func (vi *VoiceIt2) DeleteGroup(groupId string) *bytes.Buffer {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.Close()

	req, _ := http.NewRequest("DELETE", vi.BaseURL+"/groups/"+groupId, body)
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	respBody := &bytes.Buffer{}
	respBody.ReadFrom(resp.Body)

	return respBody
}

// Verify user with the given
//	userId
// (begins with 'usr_') and
//	contentLanguage
// ('en-US','es-ES', etc.)
// Note: File recording need to be no less than 1.2 seconds and no more than 5 seconds
func (vi *VoiceIt2) VoiceVerification(userId string, contentLanguage string, filePath string) *bytes.Buffer {

	file, _ := os.Open(filePath)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, _ := writer.CreateFormFile("recording", filepath.Base(file.Name()))
	io.Copy(part, file)
	_ = writer.WriteField("userId", userId)
	_ = writer.WriteField("contentLanguage", contentLanguage)
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseURL+"/verification", body)
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	respBody := &bytes.Buffer{}
	respBody.ReadFrom(resp.Body)

	return respBody
}

func (vi *VoiceIt2) VoiceVerificationByUrl(userId string, contentLanguage string, fileUrl string) *bytes.Buffer {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	_ = writer.WriteField("userId", userId)
	_ = writer.WriteField("contentLanguage", contentLanguage)
	_ = writer.WriteField("fileUrl", fileUrl)
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseURL+"/verification/byUrl", body)
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	respBody := &bytes.Buffer{}
	respBody.ReadFrom(resp.Body)

	return respBody
}

func (vi *VoiceIt2) FaceVerification(userId string, filePath string, doBlinkDetection ...bool) *bytes.Buffer {

	file, _ := os.Open(filePath)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, _ := writer.CreateFormFile("video", filepath.Base(file.Name()))
	io.Copy(part, file)
	_ = writer.WriteField("userId", userId)
	if len(doBlinkDetection) > 0 {
		_ = writer.WriteField("doBlinkDetection", strconv.FormatBool(doBlinkDetection[0]))
	}
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseURL+"/verification/face", body)
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	respBody := &bytes.Buffer{}
	respBody.ReadFrom(resp.Body)

	return respBody
}

// Verify user with given
//	userId
// (begins with 'usr_') and
//	contentLanguage
// ('en-US','es-ES', etc.)
// Note: File recording need to be no less than 1.2 seconds and no more than 5 seconds
func (vi *VoiceIt2) VideoVerification(userId string, contentLanguage string, filePath string, doBlinkDetection ...bool) *bytes.Buffer {

	file, _ := os.Open(filePath)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, _ := writer.CreateFormFile("video", filepath.Base(file.Name()))
	io.Copy(part, file)
	_ = writer.WriteField("userId", userId)
	_ = writer.WriteField("contentLanguage", contentLanguage)
	if len(doBlinkDetection) > 0 {
		_ = writer.WriteField("doBlinkDetection", strconv.FormatBool(doBlinkDetection[0]))
	}
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseURL+"/verification/video", body)
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	respBody := &bytes.Buffer{}
	respBody.ReadFrom(resp.Body)

	return respBody
}

func (vi *VoiceIt2) VideoVerificationByUrl(userId string, contentLanguage string, fileUrl string, doBlinkDetection ...bool) *bytes.Buffer {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	_ = writer.WriteField("userId", userId)
	_ = writer.WriteField("contentLanguage", contentLanguage)
	_ = writer.WriteField("fileUrl", fileUrl)
	if len(doBlinkDetection) > 0 {
		_ = writer.WriteField("doBlinkDetection", strconv.FormatBool(doBlinkDetection[0]))
	}
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseURL+"/verification/video/byUrl", body)
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	respBody := &bytes.Buffer{}
	respBody.ReadFrom(resp.Body)

	return respBody
}

// Identify user inside group with the given
//	groupId
// (begins with 'grp_') and
//	contentLanguage
// ('en-US','es-ES', etc.)
// Note: File recording need to be no less than 1.2 seconds and no more than 5 seconds
func (vi *VoiceIt2) VoiceIdentification(groupId string, contentLanguage string, filePath string) *bytes.Buffer {

	file, _ := os.Open(filePath)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, _ := writer.CreateFormFile("recording", filepath.Base(file.Name()))
	io.Copy(part, file)
	_ = writer.WriteField("groupId", groupId)
	_ = writer.WriteField("contentLanguage", contentLanguage)
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseURL+"/identification", body)
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	respBody := &bytes.Buffer{}
	respBody.ReadFrom(resp.Body)

	return respBody
}

func (vi *VoiceIt2) VoiceIdentificationByUrl(groupId string, contentLanguage string, fileUrl string) *bytes.Buffer {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	_ = writer.WriteField("fileUrl", fileUrl)
	_ = writer.WriteField("groupId", groupId)
	_ = writer.WriteField("contentLanguage", contentLanguage)
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseURL+"/identification/byUrl", body)
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	respBody := &bytes.Buffer{}
	respBody.ReadFrom(resp.Body)

	return respBody
}

// Identify user inside group with the given
//	groupId
// (begins with 'grp_') and
//	contentLanguage
// ('en-US','es-ES', etc.)
// Note: File recording need to be no less than 1.2 seconds and no more than 5 seconds
func (vi *VoiceIt2) VideoIdentification(groupId string, contentLanguage string, filePath string, doBlinkDetection ...bool) *bytes.Buffer {

	file, _ := os.Open(filePath)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, _ := writer.CreateFormFile("video", filepath.Base(file.Name()))
	io.Copy(part, file)
	_ = writer.WriteField("groupId", groupId)
	_ = writer.WriteField("contentLanguage", contentLanguage)
	if len(doBlinkDetection) > 0 {
		_ = writer.WriteField("doBlinkDetection", strconv.FormatBool(doBlinkDetection[0]))
	}
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseURL+"/identification/video", body)
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	respBody := &bytes.Buffer{}
	respBody.ReadFrom(resp.Body)

	return respBody
}

func (vi *VoiceIt2) VideoIdentificationByUrl(groupId string, contentLanguage string, fileUrl string, doBlinkDetection ...bool) *bytes.Buffer {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	_ = writer.WriteField("fileUrl", fileUrl)
	_ = writer.WriteField("groupId", groupId)
	_ = writer.WriteField("contentLanguage", contentLanguage)
	if len(doBlinkDetection) > 0 {
		_ = writer.WriteField("doBlinkDetection", strconv.FormatBool(doBlinkDetection[0]))
	}
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseURL+"/identification/video/byUrl", body)
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	respBody := &bytes.Buffer{}
	respBody.ReadFrom(resp.Body)

	return respBody
}
