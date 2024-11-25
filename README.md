# table2json
html table convert to json


# how to use
```
go get github.com/lzh-1625/table2json
```

# example
```
package main

import (
	"fmt"

	"github.com/lzh-1625/table2json"
)

func main() {
	html := `
<table class="wikitable">

<tbody><tr>
<th>Generation
</th>
<th>"A" Number
</th>
<th>Bootrom
</th>
<th>FCC ID
</th>
<th>Internal Name
</th>
<th>Identifier
</th>
<th>Finish
</th>
<th>Storage
</th>
<th>Model
</th></tr>
<tr>
<td rowspan="6"><a href="/wiki/IPad" title="IPad">iPad</a>
</td>
<td rowspan="3">A1219
</td>
<td rowspan="6"><a href="/wiki/Bootrom_574.4" title="Bootrom 574.4">Bootrom 574.4</a>
</td>
<td rowspan="3"><a rel="nofollow" class="external text" href="https://fccid.io/BCG-E2381A">BCG-E2381A</a>
</td>
<td rowspan="6">K48AP
</td>
<td rowspan="6">iPad1,1
</td>
<td rowspan="6">Black
</td>
<td>16 GB
</td>
<td>MB292
</td></tr>
<tr>
<td>32 GB
</td>
<td>MB293
</td></tr>
<tr>
<td>64 GB
</td>
<td>MB294
</td></tr>
<tr>
<td rowspan="3">A1337
</td>
<td rowspan="3"><a rel="nofollow" class="external text" href="https://fccid.io/BCG-E2328A">BCG-E2328A</a>
</td>
<td>16 GB
</td>
<td>MC349
</td></tr>
<tr>
<td>32 GB
</td>
<td>MC496
</td></tr>
<tr>
<td>64 GB
</td>
<td>MC497
</td></tr>
</tbody></table>	
	`
	result, _ := table2json.ConvertJsonString(html)
	fmt.Printf("result: %v\n", result)
}

```

## table
<table class="wikitable">

<tbody><tr>
<th>Generation
</th>
<th>"A" Number
</th>
<th>Bootrom
</th>
<th>FCC ID
</th>
<th>Internal Name
</th>
<th>Identifier
</th>
<th>Finish
</th>
<th>Storage
</th>
<th>Model
</th></tr>
<tr>
<td rowspan="6"><a href="/wiki/IPad" title="IPad">iPad</a>
</td>
<td rowspan="3">A1219
</td>
<td rowspan="6"><a href="/wiki/Bootrom_574.4" title="Bootrom 574.4">Bootrom 574.4</a>
</td>
<td rowspan="3"><a rel="nofollow" class="external text" href="https://fccid.io/BCG-E2381A">BCG-E2381A</a>
</td>
<td rowspan="6">K48AP
</td>
<td rowspan="6">iPad1,1
</td>
<td rowspan="6">Black
</td>
<td>16 GB
</td>
<td>MB292
</td></tr>
<tr>
<td>32 GB
</td>
<td>MB293
</td></tr>
<tr>
<td>64 GB
</td>
<td>MB294
</td></tr>
<tr>
<td rowspan="3">A1337
</td>
<td rowspan="3"><a rel="nofollow" class="external text" href="https://fccid.io/BCG-E2328A">BCG-E2328A</a>
</td>
<td>16 GB
</td>
<td>MC349
</td></tr>
<tr>
<td>32 GB
</td>
<td>MC496
</td></tr>
<tr>
<td>64 GB
</td>
<td>MC497
</td></tr>
</tbody></table>	

## json result
```
[
    {
        "\"A\" Number": "A1219",   
        "Bootrom": "Bootrom 574.4",
        "FCC ID": "BCG-E2381A",    
        "Finish": "Black",
        "Generation": "iPad",      
        "Identifier": "iPad1,1",   
        "Internal Name": "K48AP",  
        "Model": "MB292",
        "Storage": "16 GB"
    },
    {
        "\"A\" Number": "A1219",   
        "Bootrom": "Bootrom 574.4",
        "FCC ID": "BCG-E2381A",    
        "Finish": "Black",
        "Generation": "iPad",      
        "Identifier": "iPad1,1",   
        "Internal Name": "K48AP",
        "Model": "MB293",
        "Storage": "32 GB"
    },
    {
        "\"A\" Number": "A1219",
        "Bootrom": "Bootrom 574.4",
        "FCC ID": "BCG-E2381A",
        "Finish": "Black",
        "Generation": "iPad",
        "Identifier": "iPad1,1",
        "Internal Name": "K48AP",
        "Model": "MB294",
        "Storage": "64 GB"
    },
    {
        "\"A\" Number": "A1337",
        "Bootrom": "Bootrom 574.4",
        "FCC ID": "BCG-E2328A",
        "Finish": "Black",
        "Generation": "iPad",
        "Identifier": "iPad1,1",
        "Internal Name": "K48AP",
        "Model": "MC349",
        "Storage": "16 GB"
    },
    {
        "\"A\" Number": "A1337",
        "Bootrom": "Bootrom 574.4",
        "FCC ID": "BCG-E2328A",
        "Finish": "Black",
        "Generation": "iPad",
        "Identifier": "iPad1,1",
        "Internal Name": "K48AP",
        "Model": "MC496",
        "Storage": "32 GB"
    },
    {
        "\"A\" Number": "A1337",
        "Bootrom": "Bootrom 574.4",
        "FCC ID": "BCG-E2328A",
        "Finish": "Black",
        "Generation": "iPad",
        "Identifier": "iPad1,1",
        "Internal Name": "K48AP",
        "Model": "MC497",
        "Storage": "64 GB"
    }
]
```