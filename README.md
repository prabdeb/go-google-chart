# go-google-chart

Go utility for converting JSON data type to [Google Chart](https://google-developers.appspot.com/chart
) structure so that DataTable can understand.

This utility export the types needed for creating 
[Data Parameters](https://developers.google.com/chart/interactive/docs/reference#dataparam) needed to draw the chart

## Examples

1. Import Chart

    ```go
    import (
        GoogleChart "github.com/prabdeb/go-google-chart"
    )
    ```

2. Initiate the chart object with Columns

    ```go
    zooData := GoogleChart.Chart{
        Cols: []GoogleChart.ColsType{
            GoogleChart.ColsType{
                Label: "Animals",
                Type:  "string",
            },
            GoogleChart.ColsType{
                Label: "Count",
                Type:  "number",
            },
        },
    }
    ```

3. Insert into Rows

    ```go
    for animal, count := range zooDatabase["zoo_name"].(map[string]interface{}) {
        zooData.Rows = append(zooData.Rows, GoogleChart.RowType{
            C: []GoogleChart.RowCType{
                GoogleChart.RowCType{
                    V: animal,
                },
                GoogleChart.RowCType{
                    V: count,
                },
            },
        })
    }
    ```