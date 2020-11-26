# Burgundy
A library to take a list of things and make a report out of them.

**WARNING: Currently due to an upstream dependency, if you are going to use Google Spreadsheets you MUST use Go 1.13.x or earlier.

The core of the system is just prepping the headers and rows for the reporters. The reporters are where the real work happens. Say you have a slice of a model named `Lamp` you could generate a CSV report like this:

```go
lamps := getLamps()
data, err := burgundy.Process(lamps, CSVReporter{}); 
check(err)
err = ioutil.WriteFile("report.csv", data, 0644)
check(err)
```
