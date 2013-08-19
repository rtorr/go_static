package assert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var should_be_true string = "Should be True"

/* API */
func TestReadFileSystem(t *testing.T) {
	// Goal: is to read the file system and build an api object
	assert.True(t, false, should_be_true)
}

func TestReadFiles(t *testing.T) {
	// Goal: Get meta data and other information from
	// the pages/files to add the to the static API
	assert.True(t, false, should_be_true)
}

func TestGenerateApiFromFileSystemAndMetaData(t *testing.T) {
	// Goal Take the file system object, apply meta data from
	// pages to build the static site API
	assert.True(t, false, should_be_true)
}

/* Build */
func TestSupportFragmentHTML(t *testing.T) {
	// Goal: Assume we support html fragment files
	// which could be useful for signle page type apps

	// Point: These will be added to the API
	assert.True(t, false, should_be_true)
}

func TestTakeFilesAndBuildHtml(t *testing.T) {
	// Goal: Take content files and generate html
	assert.True(t, false, should_be_true)
}

/* Tools */
func TestGruntSupport(t *testing.T) {
	// Goal: Have built in support for task runners,
	// grunt being the first one I test

	// Point: Grunt is a wildy used tool for front end
	// developers, and that is who will be using this tool
	assert.True(t, false, should_be_true)
}

/* Misc Goals */
func TestNoPlugins(t *testing.T) {
	// Point: Plugins are a terrible idea, and no one
	// should ever make a developer use them.
	assert.True(t, true, should_be_true)
}
