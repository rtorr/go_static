package assert

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

var should_be_true string = "Should be True"

func TestReadFileSystem(t *testing.T) {
    // Goal: is to read the file system and build an api object
    assert.True(t, false, should_be_true)
}

func TestReadFiles(t *testing.T) {
    // Goal: Get meta data and other information from
    // the pages/files to add the to the static API
    assert.True(t, false, should_be_true)
}

func TestGenerateFileSystemAndMetaDataApi(t *testing.T) {
    // Goal Take the file system object, apply meta data from
    // pages to build the static site API
    assert.True(t, false, should_be_true)
}

func TestTakeFilesAndBuildHtml(t *testing.T) {
    // Goal: Take content files and generate html
    assert.True(t, false, should_be_true)
}