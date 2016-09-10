# assert
[![Build Status](https://travis-ci.org/tietang/assert.svg?branch=master)](https://travis-ci.org/tietang/assert

a testing assert lib for golang


func TestExample(t *testing.T) {
 
	assert.Equal(t, 1, 1)
  var a *A=nil
	assert.Nil(t, a)
	assert.NotNil(t, a)
 

}
