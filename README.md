This package was build to beautify Gin Gonic Binding validation error message for api. You can register your own custom validation and easily add customize its message as you like.

## Available Scripts

## Install

```bash
go get github.com/roshanr83/validator
```

## Import in your code
```
import "github.com/roshanr83/validator"
```

## Simple Examples
```
func main () {
router := gin.Default()

router.Use( validator.Errors())
	{
        router.GET("/some-method", SomeFunction)
	}

router.Run()
}
```


It has default validation message for email, required, min, max, len for extra other validation First you need to register your custom validation by process [here](https://github.com/gin-gonic/gin#custom-validators)

After you have successfully registered your custom validator you can register it when application starts in main function to beautify message by:



```
func main () {
router := gin.Default()

validate := []validator.ExtraValidation{
		{Tag: "test", Message:"%s is passed!"},
		{Tag: "unique", Message:"%s must be unique!"},
	}

validator.MakeExtraValidation(validate)

router.Use( validator.Errors())
	{
        router.GET("/some-method", SomeFunction)
	}

router.Run()
}

type TestModel struct {
    Name `json:"name" binding:required,unique`
}

func SomeFunction (c *gin.Context) {
    var model TestModel
	if err := c.ShouldBindBodyWith(&model, binding.JSON); err != nil {
		_ = c.AbortWithError(422, err).SetType(gin.ErrorTypeBind)
		return
	}
}
```


Here %s in field name which is automatically replaced later and it doesnot matter if the field name is camel case because it will be splitted automatically: for example: "firstName" will be outputted as "First name". You should name your field name with camelcase only and avoiding _ or underscore. If you have additional param like while specifying length we do len=10 and in message to get that param value just add another "%s" in message like "%s must be of length %s!" and it will be outputted like "Field name must be of length 10!"


## Test
 For testing run

 ```bash
 go test
 ```