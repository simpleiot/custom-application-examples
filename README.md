# Examples of how to use Simple IoT in your custom applications

The Simple IoT simpleiot Start() function cannot be used directly
because the frontend assets are not checked into Git. One way
to work around this is to embed the source as a submodule in your
project and then use go module replace directories.

| example                | description                           |
| ---------------------- | ------------------------------------- |
| [example 1](example-1) | Embeddeding Simple IoT as a submodule |
