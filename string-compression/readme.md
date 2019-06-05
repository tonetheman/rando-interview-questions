# notes

Only thing of note is me not remembering how to change an int to ascii in golang.
I was hoping bytes.Buffer would handle it better. There is a similar thing in Java called StringBuffer which can handle conversions like this.

I left debug junk in there but I left comments in the code where I make choices on resetting state variables.

I learned that style of changing state from a language called ReportWriter that let you process a full line at a time. So you had to keep variables around to hold state to know when to do breaks or changes for aggregation.
Old school stuff but meh.
