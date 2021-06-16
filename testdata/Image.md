<h3 id="img">Images</h3>

Admittedly, it's fairly difficult to devise a "natural" syntax for
placing images into a plain text document format.

Markdown uses an image syntax that is intended to resemble the syntax
for links, allowing for two styles: *inline* and *reference*.

Inline image syntax looks like this:

    ![Alt text](./image/fpdf.png)

    ![Alt text](./image/hiking.png "Optional title")

That is:

*   An exclamation mark: `!`;
*   followed by a set of square brackets, containing the `alt`
    attribute text for the image;
*   followed by a set of parentheses, containing the URL or path to
    the image, and an optional `title` attribute enclosed in double
    or single quotes.

Here is the first picture: ![from https://github.com/jung-kurt/gofpdf/tree/master/image](./image/fpdf.png)

Here is the second picture:
![from https://github.com/egonelbre/gophers](./image/hiking.png "Optional title")

The Go gopher was designed by Renee French. The Gopher character design is licensed under the Creative Commons 3.0 Attributions license. Read http://blog.golang.org/gopher for more details.

Note: this snippet was adapted from the testdata folder file named: 
`Markdown Documentation - Syntax.text`

Here is a non-existent image... should generate a message in trace file.
![Not from https://jpeg.org/images/jpeg-home.jpg](./image/xbay.jpg "Does not exist!")

Here is a JPEG image... is it auto-detected?
![from https://jpeg.org/images/jpeg-home.jpg](./image/bay.jpg "Down by the Bay")