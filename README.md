
# pipe-mbx
Generate Mapbox-ready layers from open source data

# Command Line 

```
go build // to build

./pipe-mbx  -t dosha-saigai-keikai-kuiki -r /Users/takutosuzuki/Downloads/A33-20_00_GML.zip -o koko.geojson
```

## Options 

| Option | Short version | Required/Default value | Comment |
|--|--|--|--|
| type | t | Yes | data source to generate - see below Supported Data Source for all available options |
| raw | r | Yes | location of raw file from source |
| output | o | No/merged.geojson | Location to output geojson file |


# Supported Data Source

| Data Type | Region | Source | 
|--|--|--|
| dosha-saigai-keikai-kuiki | Japan | [Ministry of Land, Infrastructure, Transport and Tourism (国土交通省)][1] | 


[1]: https://nlftp.mlit.go.jp/ksj/gml/datalist/KsjTmplt-A33-v1_4.html