
## Documentation

This package is using chromium for generating the PDFS, so please make sure you have it installed on your machine.

If you are on the Fedora based systems you can use following command for installation of the chromium package

#### sudo dnf install chromium


### Flexbox supported!
### Javascript supported!
### Images supported!





## Usage/Examples

This package can generate PDF by providing html or url to it.
### It takes 0.3 - 0.7 seconds for a PDF to be generated when HTML template is used (depends on complexity of the document)


## Contributing

Contributions are always welcome!

See `contributing.md` for ways to get started.

Please adhere to this project's `code of conduct`.


## Features

- Generating form HTML templates
- Generating PDF from the URL
- Waits until assets are loaded
- Custom wait (in seconds) supported


## FAQ

#### PDF is generated too early, animations did not finish

Please set waitTime, it should be greater then amount of seconds that your animations needs to load properly 

#### Can I use it with graphs?

Yes, but make sure waitTime set, since graphs sometimes need a second or two to be generated


#### Is it production stable?
No, I don't recommend using this package until alpha testing is finished.


## Authors

- [Nikola Cirkovic | @CirkaN](https://www.github.com/CirkaN)


## Feedback

If you have any feedback, please reach out to us at ncodesoft@gmail.com
