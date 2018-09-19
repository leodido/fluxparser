
//line /tmp/flux.rl:1
package fluxparser

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/influxdata/flux/ast"
)


//line /tmp/flux.rl:163



//line /tmp/flux.rl:173



//line /tmp/flux.go:21
var _actions []byte = []byte{
	0, 1, 0, 1, 1, 1, 2, 1, 3, 
	1, 4, 1, 5, 1, 6, 1, 7, 
	2, 0, 1, 2, 0, 2, 2, 0, 
	4, 2, 0, 5, 2, 1, 3, 2, 
	1, 4, 2, 1, 5, 2, 2, 0, 
	2, 3, 1, 2, 4, 0, 2, 4, 
	5, 2, 4, 6, 2, 4, 7, 3, 
	0, 1, 4, 3, 0, 4, 5, 3, 
	0, 4, 6, 3, 1, 4, 0, 3, 
	4, 0, 5, 3, 4, 0, 6, 
}

var _key_offsets []int16 = []int16{
	0, 0, 8, 20, 25, 30, 32, 34, 
	47, 60, 66, 71, 73, 86, 92, 97, 
	111, 125, 139, 153, 167, 180, 191, 205, 
	219, 233, 247, 261, 274, 288, 302, 316, 
	330, 344, 357, 368, 382, 396, 410, 424, 
	438, 451, 464, 477, 490, 503, 515, 525, 
	538, 551, 564, 577, 590, 602, 607, 612, 
	614, 616, 629, 635, 640, 654, 668, 682, 
	696, 710, 723, 734, 747, 753, 758, 772, 
	786, 800, 814, 828, 841, 854, 867, 880, 
	893, 905, 915, 927, 932, 937, 939, 941, 
	954, 960, 965, 979, 993, 1007, 1021, 1035, 
	1048, 1059, 1072, 1078, 1083, 1097, 1111, 1125, 
	1139, 1153, 1166, 1179, 1192, 1205, 1218, 1230, 
	1235, 1240, 1242, 1255, 1268, 1281, 1294, 1307, 
	1319, 1329, 1342, 1355, 1368, 1381, 1394, 1407, 
	1412, 1424, 1429, 1434, 1436, 1438, 1452, 1458, 
	1471, 1477, 1482, 1496, 1510, 1524, 1538, 1552, 
	1566, 1579, 1590, 1603, 1609, 1614, 1628, 1642, 
	1656, 1670, 1684, 1698, 1715, 1730, 1744, 1757, 
	1770, 1783, 1796, 1809, 1821, 1831, 1843, 1848, 
	1853, 1855, 1857, 1871, 1877, 1890, 1896, 1901, 
	1915, 1929, 1943, 1957, 1971, 1984, 1995, 2008, 
	2014, 2019, 2033, 2047, 2061, 2075, 2089, 2103, 
	2120, 2135, 2149, 2162, 2175, 2188, 2201, 2214, 
	2230, 2244, 2257, 2270, 2283, 2296, 2309, 2322, 
	2338, 2352, 2365, 2381, 2395, 2408, 2423, 2436, 
	2449, 2462, 2475, 2488, 2504, 2518, 2531, 2544, 
	2557, 2570, 2586, 2600, 2613, 2628, 2641, 2653, 
	2668, 2681, 2681, 2683, 
}

var _trans_keys []byte = []byte{
	95, 111, 114, 123, 65, 90, 97, 122, 
	10, 32, 61, 95, 9, 13, 48, 57, 
	65, 90, 97, 122, 10, 32, 61, 9, 
	13, 10, 32, 34, 9, 13, 10, 34, 
	10, 34, 10, 32, 34, 95, 111, 114, 
	123, 9, 13, 65, 90, 97, 122, 10, 
	32, 34, 61, 95, 9, 13, 48, 57, 
	65, 90, 97, 122, 10, 32, 34, 61, 
	9, 13, 10, 32, 34, 9, 13, 10, 
	34, 10, 32, 34, 61, 95, 9, 13, 
	48, 57, 65, 90, 97, 122, 10, 32, 
	34, 61, 9, 13, 10, 32, 34, 9, 
	13, 10, 32, 34, 61, 95, 112, 9, 
	13, 48, 57, 65, 90, 97, 122, 10, 
	32, 34, 61, 95, 116, 9, 13, 48, 
	57, 65, 90, 97, 122, 10, 32, 34, 
	61, 95, 105, 9, 13, 48, 57, 65, 
	90, 97, 122, 10, 32, 34, 61, 95, 
	111, 9, 13, 48, 57, 65, 90, 97, 
	122, 10, 32, 34, 61, 95, 110, 9, 
	13, 48, 57, 65, 90, 97, 122, 10, 
	32, 34, 61, 95, 9, 13, 48, 57, 
	65, 90, 97, 122, 10, 32, 34, 61, 
	95, 9, 13, 65, 90, 97, 122, 10, 
	32, 34, 61, 95, 101, 9, 13, 48, 
	57, 65, 90, 97, 122, 10, 32, 34, 
	61, 95, 116, 9, 13, 48, 57, 65, 
	90, 97, 122, 10, 32, 34, 61, 95, 
	117, 9, 13, 48, 57, 65, 90, 97, 
	122, 10, 32, 34, 61, 95, 114, 9, 
	13, 48, 57, 65, 90, 97, 122, 10, 
	32, 34, 61, 95, 110, 9, 13, 48, 
	57, 65, 90, 97, 122, 10, 32, 34, 
	95, 111, 114, 123, 9, 13, 65, 90, 
	97, 122, 10, 32, 34, 61, 95, 112, 
	9, 13, 48, 57, 65, 90, 97, 122, 
	10, 32, 34, 61, 95, 116, 9, 13, 
	48, 57, 65, 90, 97, 122, 10, 32, 
	34, 61, 95, 105, 9, 13, 48, 57, 
	65, 90, 97, 122, 10, 32, 34, 61, 
	95, 111, 9, 13, 48, 57, 65, 90, 
	97, 122, 10, 32, 34, 61, 95, 110, 
	9, 13, 48, 57, 65, 90, 97, 122, 
	10, 32, 34, 61, 95, 9, 13, 48, 
	57, 65, 90, 97, 122, 10, 32, 34, 
	61, 95, 9, 13, 65, 90, 97, 122, 
	10, 32, 34, 61, 95, 101, 9, 13, 
	48, 57, 65, 90, 97, 122, 10, 32, 
	34, 61, 95, 116, 9, 13, 48, 57, 
	65, 90, 97, 122, 10, 32, 34, 61, 
	95, 117, 9, 13, 48, 57, 65, 90, 
	97, 122, 10, 32, 34, 61, 95, 114, 
	9, 13, 48, 57, 65, 90, 97, 122, 
	10, 32, 34, 61, 95, 110, 9, 13, 
	48, 57, 65, 90, 97, 122, 10, 32, 
	61, 95, 112, 9, 13, 48, 57, 65, 
	90, 97, 122, 10, 32, 61, 95, 116, 
	9, 13, 48, 57, 65, 90, 97, 122, 
	10, 32, 61, 95, 105, 9, 13, 48, 
	57, 65, 90, 97, 122, 10, 32, 61, 
	95, 111, 9, 13, 48, 57, 65, 90, 
	97, 122, 10, 32, 61, 95, 110, 9, 
	13, 48, 57, 65, 90, 97, 122, 10, 
	32, 61, 95, 9, 13, 48, 57, 65, 
	90, 97, 122, 10, 32, 61, 95, 9, 
	13, 65, 90, 97, 122, 10, 32, 61, 
	95, 101, 9, 13, 48, 57, 65, 90, 
	97, 122, 10, 32, 61, 95, 116, 9, 
	13, 48, 57, 65, 90, 97, 122, 10, 
	32, 61, 95, 117, 9, 13, 48, 57, 
	65, 90, 97, 122, 10, 32, 61, 95, 
	114, 9, 13, 48, 57, 65, 90, 97, 
	122, 10, 32, 61, 95, 110, 9, 13, 
	48, 57, 65, 90, 97, 122, 10, 32, 
	61, 95, 9, 13, 48, 57, 65, 90, 
	97, 122, 10, 32, 61, 9, 13, 10, 
	32, 34, 9, 13, 10, 34, 10, 34, 
	10, 32, 34, 61, 95, 9, 13, 48, 
	57, 65, 90, 97, 122, 10, 32, 34, 
	61, 9, 13, 10, 32, 34, 9, 13, 
	10, 32, 34, 61, 95, 112, 9, 13, 
	48, 57, 65, 90, 97, 122, 10, 32, 
	34, 61, 95, 116, 9, 13, 48, 57, 
	65, 90, 97, 122, 10, 32, 34, 61, 
	95, 105, 9, 13, 48, 57, 65, 90, 
	97, 122, 10, 32, 34, 61, 95, 111, 
	9, 13, 48, 57, 65, 90, 97, 122, 
	10, 32, 34, 61, 95, 110, 9, 13, 
	48, 57, 65, 90, 97, 122, 10, 32, 
	34, 61, 95, 9, 13, 48, 57, 65, 
	90, 97, 122, 10, 32, 34, 61, 95, 
	9, 13, 65, 90, 97, 122, 10, 32, 
	34, 61, 95, 9, 13, 48, 57, 65, 
	90, 97, 122, 10, 32, 34, 61, 9, 
	13, 10, 32, 34, 9, 13, 10, 32, 
	34, 61, 95, 101, 9, 13, 48, 57, 
	65, 90, 97, 122, 10, 32, 34, 61, 
	95, 116, 9, 13, 48, 57, 65, 90, 
	97, 122, 10, 32, 34, 61, 95, 117, 
	9, 13, 48, 57, 65, 90, 97, 122, 
	10, 32, 34, 61, 95, 114, 9, 13, 
	48, 57, 65, 90, 97, 122, 10, 32, 
	34, 61, 95, 110, 9, 13, 48, 57, 
	65, 90, 97, 122, 10, 32, 61, 95, 
	112, 9, 13, 48, 57, 65, 90, 97, 
	122, 10, 32, 61, 95, 116, 9, 13, 
	48, 57, 65, 90, 97, 122, 10, 32, 
	61, 95, 105, 9, 13, 48, 57, 65, 
	90, 97, 122, 10, 32, 61, 95, 111, 
	9, 13, 48, 57, 65, 90, 97, 122, 
	10, 32, 61, 95, 110, 9, 13, 48, 
	57, 65, 90, 97, 122, 10, 32, 61, 
	95, 9, 13, 48, 57, 65, 90, 97, 
	122, 10, 32, 61, 95, 9, 13, 65, 
	90, 97, 122, 10, 32, 61, 95, 9, 
	13, 48, 57, 65, 90, 97, 122, 10, 
	32, 61, 9, 13, 10, 32, 34, 9, 
	13, 10, 34, 10, 34, 10, 32, 34, 
	61, 95, 9, 13, 48, 57, 65, 90, 
	97, 122, 10, 32, 34, 61, 9, 13, 
	10, 32, 34, 9, 13, 10, 32, 34, 
	61, 95, 112, 9, 13, 48, 57, 65, 
	90, 97, 122, 10, 32, 34, 61, 95, 
	116, 9, 13, 48, 57, 65, 90, 97, 
	122, 10, 32, 34, 61, 95, 105, 9, 
	13, 48, 57, 65, 90, 97, 122, 10, 
	32, 34, 61, 95, 111, 9, 13, 48, 
	57, 65, 90, 97, 122, 10, 32, 34, 
	61, 95, 110, 9, 13, 48, 57, 65, 
	90, 97, 122, 10, 32, 34, 61, 95, 
	9, 13, 48, 57, 65, 90, 97, 122, 
	10, 32, 34, 61, 95, 9, 13, 65, 
	90, 97, 122, 10, 32, 34, 61, 95, 
	9, 13, 48, 57, 65, 90, 97, 122, 
	10, 32, 34, 61, 9, 13, 10, 32, 
	34, 9, 13, 10, 32, 34, 61, 95, 
	101, 9, 13, 48, 57, 65, 90, 97, 
	122, 10, 32, 34, 61, 95, 116, 9, 
	13, 48, 57, 65, 90, 97, 122, 10, 
	32, 34, 61, 95, 117, 9, 13, 48, 
	57, 65, 90, 97, 122, 10, 32, 34, 
	61, 95, 114, 9, 13, 48, 57, 65, 
	90, 97, 122, 10, 32, 34, 61, 95, 
	110, 9, 13, 48, 57, 65, 90, 97, 
	122, 10, 32, 61, 95, 101, 9, 13, 
	48, 57, 65, 90, 97, 122, 10, 32, 
	61, 95, 116, 9, 13, 48, 57, 65, 
	90, 97, 122, 10, 32, 61, 95, 117, 
	9, 13, 48, 57, 65, 90, 97, 122, 
	10, 32, 61, 95, 114, 9, 13, 48, 
	57, 65, 90, 97, 122, 10, 32, 61, 
	95, 110, 9, 13, 48, 57, 65, 90, 
	97, 122, 10, 32, 61, 95, 9, 13, 
	48, 57, 65, 90, 97, 122, 10, 32, 
	61, 9, 13, 10, 32, 34, 9, 13, 
	10, 34, 10, 32, 61, 95, 112, 9, 
	13, 48, 57, 65, 90, 97, 122, 10, 
	32, 61, 95, 116, 9, 13, 48, 57, 
	65, 90, 97, 122, 10, 32, 61, 95, 
	105, 9, 13, 48, 57, 65, 90, 97, 
	122, 10, 32, 61, 95, 111, 9, 13, 
	48, 57, 65, 90, 97, 122, 10, 32, 
	61, 95, 110, 9, 13, 48, 57, 65, 
	90, 97, 122, 10, 32, 61, 95, 9, 
	13, 48, 57, 65, 90, 97, 122, 10, 
	32, 61, 95, 9, 13, 65, 90, 97, 
	122, 10, 32, 61, 95, 101, 9, 13, 
	48, 57, 65, 90, 97, 122, 10, 32, 
	61, 95, 116, 9, 13, 48, 57, 65, 
	90, 97, 122, 10, 32, 61, 95, 117, 
	9, 13, 48, 57, 65, 90, 97, 122, 
	10, 32, 61, 95, 114, 9, 13, 48, 
	57, 65, 90, 97, 122, 10, 32, 61, 
	95, 110, 9, 13, 48, 57, 65, 90, 
	97, 122, 10, 32, 95, 111, 114, 123, 
	125, 9, 13, 65, 90, 97, 122, 10, 
	32, 125, 9, 13, 10, 32, 61, 95, 
	9, 13, 48, 57, 65, 90, 97, 122, 
	10, 32, 61, 9, 13, 10, 32, 34, 
	9, 13, 10, 34, 10, 34, 10, 32, 
	34, 95, 111, 114, 123, 125, 9, 13, 
	65, 90, 97, 122, 10, 32, 34, 125, 
	9, 13, 10, 32, 34, 61, 95, 9, 
	13, 48, 57, 65, 90, 97, 122, 10, 
	32, 34, 61, 9, 13, 10, 32, 34, 
	9, 13, 10, 32, 34, 95, 111, 114, 
	123, 125, 9, 13, 65, 90, 97, 122, 
	10, 32, 34, 61, 95, 112, 9, 13, 
	48, 57, 65, 90, 97, 122, 10, 32, 
	34, 61, 95, 116, 9, 13, 48, 57, 
	65, 90, 97, 122, 10, 32, 34, 61, 
	95, 105, 9, 13, 48, 57, 65, 90, 
	97, 122, 10, 32, 34, 61, 95, 111, 
	9, 13, 48, 57, 65, 90, 97, 122, 
	10, 32, 34, 61, 95, 110, 9, 13, 
	48, 57, 65, 90, 97, 122, 10, 32, 
	34, 61, 95, 9, 13, 48, 57, 65, 
	90, 97, 122, 10, 32, 34, 61, 95, 
	9, 13, 65, 90, 97, 122, 10, 32, 
	34, 61, 95, 9, 13, 48, 57, 65, 
	90, 97, 122, 10, 32, 34, 61, 9, 
	13, 10, 32, 34, 9, 13, 10, 32, 
	34, 95, 111, 114, 123, 125, 9, 13, 
	65, 90, 97, 122, 10, 32, 34, 61, 
	95, 101, 9, 13, 48, 57, 65, 90, 
	97, 122, 10, 32, 34, 61, 95, 116, 
	9, 13, 48, 57, 65, 90, 97, 122, 
	10, 32, 34, 61, 95, 117, 9, 13, 
	48, 57, 65, 90, 97, 122, 10, 32, 
	34, 61, 95, 114, 9, 13, 48, 57, 
	65, 90, 97, 122, 10, 32, 34, 61, 
	95, 110, 9, 13, 48, 57, 65, 90, 
	97, 122, 10, 32, 34, 61, 95, 111, 
	114, 123, 125, 9, 13, 48, 57, 65, 
	90, 97, 122, 10, 32, 34, 61, 95, 
	111, 114, 123, 125, 9, 13, 65, 90, 
	97, 122, 10, 32, 34, 95, 111, 114, 
	123, 125, 9, 13, 65, 90, 97, 122, 
	10, 32, 61, 95, 112, 9, 13, 48, 
	57, 65, 90, 97, 122, 10, 32, 61, 
	95, 116, 9, 13, 48, 57, 65, 90, 
	97, 122, 10, 32, 61, 95, 105, 9, 
	13, 48, 57, 65, 90, 97, 122, 10, 
	32, 61, 95, 111, 9, 13, 48, 57, 
	65, 90, 97, 122, 10, 32, 61, 95, 
	110, 9, 13, 48, 57, 65, 90, 97, 
	122, 10, 32, 61, 95, 9, 13, 48, 
	57, 65, 90, 97, 122, 10, 32, 61, 
	95, 9, 13, 65, 90, 97, 122, 10, 
	32, 61, 95, 9, 13, 48, 57, 65, 
	90, 97, 122, 10, 32, 61, 9, 13, 
	10, 32, 34, 9, 13, 10, 34, 10, 
	34, 10, 32, 34, 95, 111, 114, 123, 
	125, 9, 13, 65, 90, 97, 122, 10, 
	32, 34, 125, 9, 13, 10, 32, 34, 
	61, 95, 9, 13, 48, 57, 65, 90, 
	97, 122, 10, 32, 34, 61, 9, 13, 
	10, 32, 34, 9, 13, 10, 32, 34, 
	61, 95, 112, 9, 13, 48, 57, 65, 
	90, 97, 122, 10, 32, 34, 61, 95, 
	116, 9, 13, 48, 57, 65, 90, 97, 
	122, 10, 32, 34, 61, 95, 105, 9, 
	13, 48, 57, 65, 90, 97, 122, 10, 
	32, 34, 61, 95, 111, 9, 13, 48, 
	57, 65, 90, 97, 122, 10, 32, 34, 
	61, 95, 110, 9, 13, 48, 57, 65, 
	90, 97, 122, 10, 32, 34, 61, 95, 
	9, 13, 48, 57, 65, 90, 97, 122, 
	10, 32, 34, 61, 95, 9, 13, 65, 
	90, 97, 122, 10, 32, 34, 61, 95, 
	9, 13, 48, 57, 65, 90, 97, 122, 
	10, 32, 34, 61, 9, 13, 10, 32, 
	34, 9, 13, 10, 32, 34, 95, 111, 
	114, 123, 125, 9, 13, 65, 90, 97, 
	122, 10, 32, 34, 61, 95, 101, 9, 
	13, 48, 57, 65, 90, 97, 122, 10, 
	32, 34, 61, 95, 116, 9, 13, 48, 
	57, 65, 90, 97, 122, 10, 32, 34, 
	61, 95, 117, 9, 13, 48, 57, 65, 
	90, 97, 122, 10, 32, 34, 61, 95, 
	114, 9, 13, 48, 57, 65, 90, 97, 
	122, 10, 32, 34, 61, 95, 110, 9, 
	13, 48, 57, 65, 90, 97, 122, 10, 
	32, 34, 61, 95, 111, 114, 123, 125, 
	9, 13, 48, 57, 65, 90, 97, 122, 
	10, 32, 34, 61, 95, 111, 114, 123, 
	125, 9, 13, 65, 90, 97, 122, 10, 
	32, 34, 95, 111, 114, 123, 125, 9, 
	13, 65, 90, 97, 122, 10, 32, 61, 
	95, 101, 9, 13, 48, 57, 65, 90, 
	97, 122, 10, 32, 61, 95, 116, 9, 
	13, 48, 57, 65, 90, 97, 122, 10, 
	32, 61, 95, 117, 9, 13, 48, 57, 
	65, 90, 97, 122, 10, 32, 61, 95, 
	114, 9, 13, 48, 57, 65, 90, 97, 
	122, 10, 32, 61, 95, 110, 9, 13, 
	48, 57, 65, 90, 97, 122, 10, 32, 
	61, 95, 111, 114, 123, 125, 9, 13, 
	48, 57, 65, 90, 97, 122, 10, 32, 
	61, 95, 111, 114, 123, 125, 9, 13, 
	65, 90, 97, 122, 10, 32, 95, 111, 
	114, 123, 125, 9, 13, 65, 90, 97, 
	122, 10, 32, 34, 95, 111, 114, 123, 
	9, 13, 65, 90, 97, 122, 10, 32, 
	34, 95, 111, 114, 123, 9, 13, 65, 
	90, 97, 122, 10, 32, 34, 95, 111, 
	114, 123, 9, 13, 65, 90, 97, 122, 
	10, 32, 34, 95, 111, 114, 123, 9, 
	13, 65, 90, 97, 122, 10, 32, 34, 
	95, 111, 114, 123, 9, 13, 65, 90, 
	97, 122, 10, 32, 34, 61, 95, 111, 
	114, 123, 9, 13, 48, 57, 65, 90, 
	97, 122, 10, 32, 34, 61, 95, 111, 
	114, 123, 9, 13, 65, 90, 97, 122, 
	10, 32, 34, 95, 111, 114, 123, 9, 
	13, 65, 90, 97, 122, 10, 32, 34, 
	61, 95, 111, 114, 123, 9, 13, 48, 
	57, 65, 90, 97, 122, 10, 32, 34, 
	61, 95, 111, 114, 123, 9, 13, 65, 
	90, 97, 122, 10, 32, 34, 95, 111, 
	114, 123, 9, 13, 65, 90, 97, 122, 
	10, 32, 61, 95, 111, 114, 123, 9, 
	13, 48, 57, 65, 90, 97, 122, 10, 
	32, 61, 95, 111, 114, 123, 9, 13, 
	65, 90, 97, 122, 10, 32, 34, 95, 
	111, 114, 123, 9, 13, 65, 90, 97, 
	122, 10, 32, 34, 95, 111, 114, 123, 
	9, 13, 65, 90, 97, 122, 10, 32, 
	34, 95, 111, 114, 123, 9, 13, 65, 
	90, 97, 122, 10, 32, 34, 95, 111, 
	114, 123, 9, 13, 65, 90, 97, 122, 
	10, 32, 34, 61, 95, 111, 114, 123, 
	9, 13, 48, 57, 65, 90, 97, 122, 
	10, 32, 34, 61, 95, 111, 114, 123, 
	9, 13, 65, 90, 97, 122, 10, 32, 
	34, 95, 111, 114, 123, 9, 13, 65, 
	90, 97, 122, 10, 32, 34, 95, 111, 
	114, 123, 9, 13, 65, 90, 97, 122, 
	10, 32, 34, 95, 111, 114, 123, 9, 
	13, 65, 90, 97, 122, 10, 32, 34, 
	95, 111, 114, 123, 9, 13, 65, 90, 
	97, 122, 10, 32, 34, 61, 95, 111, 
	114, 123, 9, 13, 48, 57, 65, 90, 
	97, 122, 10, 32, 34, 61, 95, 111, 
	114, 123, 9, 13, 65, 90, 97, 122, 
	10, 32, 34, 95, 111, 114, 123, 9, 
	13, 65, 90, 97, 122, 10, 32, 61, 
	95, 111, 114, 123, 9, 13, 48, 57, 
	65, 90, 97, 122, 10, 32, 61, 95, 
	111, 114, 123, 9, 13, 65, 90, 97, 
	122, 10, 32, 95, 111, 114, 123, 9, 
	13, 65, 90, 97, 122, 10, 32, 61, 
	95, 111, 114, 123, 9, 13, 48, 57, 
	65, 90, 97, 122, 10, 32, 61, 95, 
	111, 114, 123, 9, 13, 65, 90, 97, 
	122, 10, 34, 10, 34, 
}

var _single_lengths []byte = []byte{
	0, 4, 4, 3, 3, 2, 2, 7, 
	5, 4, 3, 2, 5, 4, 3, 6, 
	6, 6, 6, 6, 5, 5, 6, 6, 
	6, 6, 6, 7, 6, 6, 6, 6, 
	6, 5, 5, 6, 6, 6, 6, 6, 
	5, 5, 5, 5, 5, 4, 4, 5, 
	5, 5, 5, 5, 4, 3, 3, 2, 
	2, 5, 4, 3, 6, 6, 6, 6, 
	6, 5, 5, 5, 4, 3, 6, 6, 
	6, 6, 6, 5, 5, 5, 5, 5, 
	4, 4, 4, 3, 3, 2, 2, 5, 
	4, 3, 6, 6, 6, 6, 6, 5, 
	5, 5, 4, 3, 6, 6, 6, 6, 
	6, 5, 5, 5, 5, 5, 4, 3, 
	3, 2, 5, 5, 5, 5, 5, 4, 
	4, 5, 5, 5, 5, 5, 7, 3, 
	4, 3, 3, 2, 2, 8, 4, 5, 
	4, 3, 8, 6, 6, 6, 6, 6, 
	5, 5, 5, 4, 3, 8, 6, 6, 
	6, 6, 6, 9, 9, 8, 5, 5, 
	5, 5, 5, 4, 4, 4, 3, 3, 
	2, 2, 8, 4, 5, 4, 3, 6, 
	6, 6, 6, 6, 5, 5, 5, 4, 
	3, 8, 6, 6, 6, 6, 6, 9, 
	9, 8, 5, 5, 5, 5, 5, 8, 
	8, 7, 7, 7, 7, 7, 7, 8, 
	8, 7, 8, 8, 7, 7, 7, 7, 
	7, 7, 7, 8, 8, 7, 7, 7, 
	7, 8, 8, 7, 7, 7, 6, 7, 
	7, 0, 2, 2, 
}

var _range_lengths []byte = []byte{
	0, 2, 4, 1, 1, 0, 0, 3, 
	4, 1, 1, 0, 4, 1, 1, 4, 
	4, 4, 4, 4, 4, 3, 4, 4, 
	4, 4, 4, 3, 4, 4, 4, 4, 
	4, 4, 3, 4, 4, 4, 4, 4, 
	4, 4, 4, 4, 4, 4, 3, 4, 
	4, 4, 4, 4, 4, 1, 1, 0, 
	0, 4, 1, 1, 4, 4, 4, 4, 
	4, 4, 3, 4, 1, 1, 4, 4, 
	4, 4, 4, 4, 4, 4, 4, 4, 
	4, 3, 4, 1, 1, 0, 0, 4, 
	1, 1, 4, 4, 4, 4, 4, 4, 
	3, 4, 1, 1, 4, 4, 4, 4, 
	4, 4, 4, 4, 4, 4, 4, 1, 
	1, 0, 4, 4, 4, 4, 4, 4, 
	3, 4, 4, 4, 4, 4, 3, 1, 
	4, 1, 1, 0, 0, 3, 1, 4, 
	1, 1, 3, 4, 4, 4, 4, 4, 
	4, 3, 4, 1, 1, 3, 4, 4, 
	4, 4, 4, 4, 3, 3, 4, 4, 
	4, 4, 4, 4, 3, 4, 1, 1, 
	0, 0, 3, 1, 4, 1, 1, 4, 
	4, 4, 4, 4, 4, 3, 4, 1, 
	1, 3, 4, 4, 4, 4, 4, 4, 
	3, 3, 4, 4, 4, 4, 4, 4, 
	3, 3, 3, 3, 3, 3, 3, 4, 
	3, 3, 4, 3, 3, 4, 3, 3, 
	3, 3, 3, 4, 3, 3, 3, 3, 
	3, 4, 3, 3, 4, 3, 3, 4, 
	3, 0, 0, 0, 
}

var _index_offsets []int16 = []int16{
	0, 0, 7, 16, 21, 26, 29, 32, 
	43, 53, 59, 64, 67, 77, 83, 88, 
	99, 110, 121, 132, 143, 153, 162, 173, 
	184, 195, 206, 217, 228, 239, 250, 261, 
	272, 283, 293, 302, 313, 324, 335, 346, 
	357, 367, 377, 387, 397, 407, 416, 424, 
	434, 444, 454, 464, 474, 483, 488, 493, 
	496, 499, 509, 515, 520, 531, 542, 553, 
	564, 575, 585, 594, 604, 610, 615, 626, 
	637, 648, 659, 670, 680, 690, 700, 710, 
	720, 729, 737, 746, 751, 756, 759, 762, 
	772, 778, 783, 794, 805, 816, 827, 838, 
	848, 857, 867, 873, 878, 889, 900, 911, 
	922, 933, 943, 953, 963, 973, 983, 992, 
	997, 1002, 1005, 1015, 1025, 1035, 1045, 1055, 
	1064, 1072, 1082, 1092, 1102, 1112, 1122, 1133, 
	1138, 1147, 1152, 1157, 1160, 1163, 1175, 1181, 
	1191, 1197, 1202, 1214, 1225, 1236, 1247, 1258, 
	1269, 1279, 1288, 1298, 1304, 1309, 1321, 1332, 
	1343, 1354, 1365, 1376, 1390, 1403, 1415, 1425, 
	1435, 1445, 1455, 1465, 1474, 1482, 1491, 1496, 
	1501, 1504, 1507, 1519, 1525, 1535, 1541, 1546, 
	1557, 1568, 1579, 1590, 1601, 1611, 1620, 1630, 
	1636, 1641, 1653, 1664, 1675, 1686, 1697, 1708, 
	1722, 1735, 1747, 1757, 1767, 1777, 1787, 1797, 
	1810, 1822, 1833, 1844, 1855, 1866, 1877, 1888, 
	1901, 1913, 1924, 1937, 1949, 1960, 1972, 1983, 
	1994, 2005, 2016, 2027, 2040, 2052, 2063, 2074, 
	2085, 2096, 2109, 2121, 2132, 2144, 2155, 2165, 
	2177, 2188, 2189, 2192, 
}

var _trans_targs []byte = []byte{
	2, 40, 47, 230, 2, 2, 0, 3, 
	3, 4, 2, 3, 2, 2, 2, 0, 
	3, 3, 4, 3, 0, 4, 4, 5, 
	4, 0, 6, 202, 6, 6, 202, 6, 
	7, 7, 202, 8, 28, 35, 212, 7, 
	8, 8, 6, 9, 9, 202, 10, 8, 
	9, 8, 8, 8, 6, 9, 9, 202, 
	10, 9, 6, 10, 10, 203, 10, 6, 
	11, 204, 11, 13, 13, 204, 14, 12, 
	13, 12, 12, 12, 11, 13, 13, 204, 
	14, 13, 11, 14, 14, 206, 14, 11, 
	13, 13, 204, 14, 12, 16, 13, 12, 
	12, 12, 11, 13, 13, 204, 14, 12, 
	17, 13, 12, 12, 12, 11, 13, 13, 
	204, 14, 12, 18, 13, 12, 12, 12, 
	11, 13, 13, 204, 14, 12, 19, 13, 
	12, 12, 12, 11, 13, 13, 204, 14, 
	12, 20, 13, 12, 12, 12, 11, 21, 
	21, 204, 14, 12, 21, 12, 12, 12, 
	11, 21, 21, 204, 14, 12, 21, 12, 
	12, 11, 13, 13, 204, 14, 12, 23, 
	13, 12, 12, 12, 11, 13, 13, 204, 
	14, 12, 24, 13, 12, 12, 12, 11, 
	13, 13, 204, 14, 12, 25, 13, 12, 
	12, 12, 11, 13, 13, 204, 14, 12, 
	26, 13, 12, 12, 12, 11, 13, 13, 
	204, 14, 12, 207, 13, 12, 12, 12, 
	11, 27, 27, 204, 12, 15, 22, 209, 
	27, 12, 12, 11, 9, 9, 202, 10, 
	8, 29, 9, 8, 8, 8, 6, 9, 
	9, 202, 10, 8, 30, 9, 8, 8, 
	8, 6, 9, 9, 202, 10, 8, 31, 
	9, 8, 8, 8, 6, 9, 9, 202, 
	10, 8, 32, 9, 8, 8, 8, 6, 
	9, 9, 202, 10, 8, 33, 9, 8, 
	8, 8, 6, 34, 34, 202, 10, 8, 
	34, 8, 8, 8, 6, 34, 34, 202, 
	10, 8, 34, 8, 8, 6, 9, 9, 
	202, 10, 8, 36, 9, 8, 8, 8, 
	6, 9, 9, 202, 10, 8, 37, 9, 
	8, 8, 8, 6, 9, 9, 202, 10, 
	8, 38, 9, 8, 8, 8, 6, 9, 
	9, 202, 10, 8, 39, 9, 8, 8, 
	8, 6, 9, 9, 202, 10, 8, 210, 
	9, 8, 8, 8, 6, 3, 3, 4, 
	2, 41, 3, 2, 2, 2, 0, 3, 
	3, 4, 2, 42, 3, 2, 2, 2, 
	0, 3, 3, 4, 2, 43, 3, 2, 
	2, 2, 0, 3, 3, 4, 2, 44, 
	3, 2, 2, 2, 0, 3, 3, 4, 
	2, 45, 3, 2, 2, 2, 0, 46, 
	46, 4, 2, 46, 2, 2, 2, 0, 
	46, 46, 4, 2, 46, 2, 2, 0, 
	3, 3, 4, 2, 48, 3, 2, 2, 
	2, 0, 3, 3, 4, 2, 49, 3, 
	2, 2, 2, 0, 3, 3, 4, 2, 
	50, 3, 2, 2, 2, 0, 3, 3, 
	4, 2, 51, 3, 2, 2, 2, 0, 
	3, 3, 4, 2, 213, 3, 2, 2, 
	2, 0, 53, 53, 54, 52, 53, 52, 
	52, 52, 0, 53, 53, 54, 53, 0, 
	54, 54, 55, 54, 0, 56, 215, 56, 
	56, 215, 56, 58, 58, 215, 59, 57, 
	58, 57, 57, 57, 56, 58, 58, 215, 
	59, 58, 56, 59, 59, 217, 59, 56, 
	58, 58, 215, 59, 57, 61, 58, 57, 
	57, 57, 56, 58, 58, 215, 59, 57, 
	62, 58, 57, 57, 57, 56, 58, 58, 
	215, 59, 57, 63, 58, 57, 57, 57, 
	56, 58, 58, 215, 59, 57, 64, 58, 
	57, 57, 57, 56, 58, 58, 215, 59, 
	57, 65, 58, 57, 57, 57, 56, 66, 
	66, 215, 59, 57, 66, 57, 57, 57, 
	56, 66, 66, 215, 59, 67, 66, 67, 
	67, 56, 68, 68, 215, 69, 67, 68, 
	67, 67, 67, 56, 68, 68, 215, 69, 
	68, 56, 69, 69, 218, 69, 56, 58, 
	58, 215, 59, 57, 71, 58, 57, 57, 
	57, 56, 58, 58, 215, 59, 57, 72, 
	58, 57, 57, 57, 56, 58, 58, 215, 
	59, 57, 73, 58, 57, 57, 57, 56, 
	58, 58, 215, 59, 57, 74, 58, 57, 
	57, 57, 56, 58, 58, 215, 59, 57, 
	219, 58, 57, 57, 57, 56, 53, 53, 
	54, 52, 76, 53, 52, 52, 52, 0, 
	53, 53, 54, 52, 77, 53, 52, 52, 
	52, 0, 53, 53, 54, 52, 78, 53, 
	52, 52, 52, 0, 53, 53, 54, 52, 
	79, 53, 52, 52, 52, 0, 53, 53, 
	54, 52, 80, 53, 52, 52, 52, 0, 
	81, 81, 54, 52, 81, 52, 52, 52, 
	0, 81, 81, 54, 82, 81, 82, 82, 
	0, 83, 83, 84, 82, 83, 82, 82, 
	82, 0, 83, 83, 84, 83, 0, 84, 
	84, 85, 84, 0, 86, 222, 86, 86, 
	222, 86, 88, 88, 222, 89, 87, 88, 
	87, 87, 87, 86, 88, 88, 222, 89, 
	88, 86, 89, 89, 217, 89, 86, 88, 
	88, 222, 89, 87, 91, 88, 87, 87, 
	87, 86, 88, 88, 222, 89, 87, 92, 
	88, 87, 87, 87, 86, 88, 88, 222, 
	89, 87, 93, 88, 87, 87, 87, 86, 
	88, 88, 222, 89, 87, 94, 88, 87, 
	87, 87, 86, 88, 88, 222, 89, 87, 
	95, 88, 87, 87, 87, 86, 96, 96, 
	222, 89, 87, 96, 87, 87, 87, 86, 
	96, 96, 222, 89, 97, 96, 97, 97, 
	86, 98, 98, 222, 99, 97, 98, 97, 
	97, 97, 86, 98, 98, 222, 99, 98, 
	86, 99, 99, 224, 99, 86, 88, 88, 
	222, 89, 87, 101, 88, 87, 87, 87, 
	86, 88, 88, 222, 89, 87, 102, 88, 
	87, 87, 87, 86, 88, 88, 222, 89, 
	87, 103, 88, 87, 87, 87, 86, 88, 
	88, 222, 89, 87, 104, 88, 87, 87, 
	87, 86, 88, 88, 222, 89, 87, 225, 
	88, 87, 87, 87, 86, 53, 53, 54, 
	52, 106, 53, 52, 52, 52, 0, 53, 
	53, 54, 52, 107, 53, 52, 52, 52, 
	0, 53, 53, 54, 52, 108, 53, 52, 
	52, 52, 0, 53, 53, 54, 52, 109, 
	53, 52, 52, 52, 0, 53, 53, 54, 
	52, 228, 53, 52, 52, 52, 0, 111, 
	111, 112, 110, 111, 110, 110, 110, 0, 
	111, 111, 112, 111, 0, 112, 112, 113, 
	112, 0, 11, 204, 11, 111, 111, 112, 
	110, 115, 111, 110, 110, 110, 0, 111, 
	111, 112, 110, 116, 111, 110, 110, 110, 
	0, 111, 111, 112, 110, 117, 111, 110, 
	110, 110, 0, 111, 111, 112, 110, 118, 
	111, 110, 110, 110, 0, 111, 111, 112, 
	110, 119, 111, 110, 110, 110, 0, 120, 
	120, 112, 110, 120, 110, 110, 110, 0, 
	120, 120, 112, 82, 120, 82, 82, 0, 
	111, 111, 112, 110, 122, 111, 110, 110, 
	110, 0, 111, 111, 112, 110, 123, 111, 
	110, 110, 110, 0, 111, 111, 112, 110, 
	124, 111, 110, 110, 110, 0, 111, 111, 
	112, 110, 125, 111, 110, 110, 110, 0, 
	111, 111, 112, 110, 231, 111, 110, 110, 
	110, 0, 127, 127, 128, 158, 194, 201, 
	233, 127, 128, 128, 0, 127, 127, 233, 
	127, 0, 129, 129, 130, 128, 129, 128, 
	128, 128, 0, 129, 129, 130, 129, 0, 
	130, 130, 131, 130, 0, 132, 133, 132, 
	132, 133, 132, 134, 134, 133, 135, 139, 
	150, 157, 234, 134, 135, 135, 132, 134, 
	134, 133, 234, 134, 132, 136, 136, 133, 
	137, 135, 136, 135, 135, 135, 132, 136, 
	136, 133, 137, 136, 132, 137, 137, 138, 
	137, 132, 134, 134, 133, 135, 139, 150, 
	157, 234, 134, 135, 135, 132, 136, 136, 
	133, 137, 135, 140, 136, 135, 135, 135, 
	132, 136, 136, 133, 137, 135, 141, 136, 
	135, 135, 135, 132, 136, 136, 133, 137, 
	135, 142, 136, 135, 135, 135, 132, 136, 
	136, 133, 137, 135, 143, 136, 135, 135, 
	135, 132, 136, 136, 133, 137, 135, 144, 
	136, 135, 135, 135, 132, 145, 145, 133, 
	137, 135, 145, 135, 135, 135, 132, 145, 
	145, 133, 137, 146, 145, 146, 146, 132, 
	147, 147, 133, 148, 146, 147, 146, 146, 
	146, 132, 147, 147, 133, 148, 147, 132, 
	148, 148, 149, 148, 132, 134, 134, 133, 
	135, 139, 150, 157, 234, 134, 135, 135, 
	132, 136, 136, 133, 137, 135, 151, 136, 
	135, 135, 135, 132, 136, 136, 133, 137, 
	135, 152, 136, 135, 135, 135, 132, 136, 
	136, 133, 137, 135, 153, 136, 135, 135, 
	135, 132, 136, 136, 133, 137, 135, 154, 
	136, 135, 135, 135, 132, 136, 136, 133, 
	137, 135, 155, 136, 135, 135, 135, 132, 
	156, 156, 133, 137, 135, 139, 150, 157, 
	234, 156, 135, 135, 135, 132, 156, 156, 
	133, 137, 135, 139, 150, 157, 234, 156, 
	135, 135, 132, 157, 157, 133, 135, 139, 
	150, 157, 234, 157, 135, 135, 132, 129, 
	129, 130, 128, 159, 129, 128, 128, 128, 
	0, 129, 129, 130, 128, 160, 129, 128, 
	128, 128, 0, 129, 129, 130, 128, 161, 
	129, 128, 128, 128, 0, 129, 129, 130, 
	128, 162, 129, 128, 128, 128, 0, 129, 
	129, 130, 128, 163, 129, 128, 128, 128, 
	0, 164, 164, 130, 128, 164, 128, 128, 
	128, 0, 164, 164, 130, 165, 164, 165, 
	165, 0, 166, 166, 167, 165, 166, 165, 
	165, 165, 0, 166, 166, 167, 166, 0, 
	167, 167, 168, 167, 0, 169, 170, 169, 
	169, 170, 169, 171, 171, 170, 172, 175, 
	186, 193, 235, 171, 172, 172, 169, 171, 
	171, 170, 235, 171, 169, 173, 173, 170, 
	174, 172, 173, 172, 172, 172, 169, 173, 
	173, 170, 174, 173, 169, 174, 174, 138, 
	174, 169, 173, 173, 170, 174, 172, 176, 
	173, 172, 172, 172, 169, 173, 173, 170, 
	174, 172, 177, 173, 172, 172, 172, 169, 
	173, 173, 170, 174, 172, 178, 173, 172, 
	172, 172, 169, 173, 173, 170, 174, 172, 
	179, 173, 172, 172, 172, 169, 173, 173, 
	170, 174, 172, 180, 173, 172, 172, 172, 
	169, 181, 181, 170, 174, 172, 181, 172, 
	172, 172, 169, 181, 181, 170, 174, 182, 
	181, 182, 182, 169, 183, 183, 170, 184, 
	182, 183, 182, 182, 182, 169, 183, 183, 
	170, 184, 183, 169, 184, 184, 185, 184, 
	169, 171, 171, 170, 172, 175, 186, 193, 
	235, 171, 172, 172, 169, 173, 173, 170, 
	174, 172, 187, 173, 172, 172, 172, 169, 
	173, 173, 170, 174, 172, 188, 173, 172, 
	172, 172, 169, 173, 173, 170, 174, 172, 
	189, 173, 172, 172, 172, 169, 173, 173, 
	170, 174, 172, 190, 173, 172, 172, 172, 
	169, 173, 173, 170, 174, 172, 191, 173, 
	172, 172, 172, 169, 192, 192, 170, 174, 
	172, 175, 186, 193, 235, 192, 172, 172, 
	172, 169, 192, 192, 170, 174, 172, 175, 
	186, 193, 235, 192, 172, 172, 169, 193, 
	193, 170, 172, 175, 186, 193, 235, 193, 
	172, 172, 169, 129, 129, 130, 128, 195, 
	129, 128, 128, 128, 0, 129, 129, 130, 
	128, 196, 129, 128, 128, 128, 0, 129, 
	129, 130, 128, 197, 129, 128, 128, 128, 
	0, 129, 129, 130, 128, 198, 129, 128, 
	128, 128, 0, 129, 129, 130, 128, 199, 
	129, 128, 128, 128, 0, 200, 200, 130, 
	128, 158, 194, 201, 233, 200, 128, 128, 
	128, 0, 200, 200, 130, 128, 158, 194, 
	201, 233, 200, 128, 128, 0, 201, 201, 
	128, 158, 194, 201, 233, 201, 128, 128, 
	0, 7, 7, 202, 8, 28, 35, 212, 
	7, 8, 8, 6, 27, 27, 204, 12, 
	15, 22, 209, 27, 12, 12, 11, 205, 
	205, 204, 12, 15, 22, 209, 205, 12, 
	12, 11, 205, 205, 204, 12, 15, 22, 
	209, 205, 12, 12, 11, 205, 205, 204, 
	12, 15, 22, 209, 205, 12, 12, 11, 
	208, 208, 204, 14, 12, 15, 22, 209, 
	208, 12, 12, 12, 11, 208, 208, 204, 
	14, 12, 15, 22, 209, 208, 12, 12, 
	11, 209, 209, 204, 12, 15, 22, 209, 
	209, 12, 12, 11, 211, 211, 202, 10, 
	8, 28, 35, 212, 211, 8, 8, 8, 
	6, 211, 211, 202, 10, 8, 28, 35, 
	212, 211, 8, 8, 6, 212, 212, 202, 
	8, 28, 35, 212, 212, 8, 8, 6, 
	214, 214, 4, 110, 114, 121, 230, 214, 
	2, 110, 110, 0, 214, 214, 4, 52, 
	75, 105, 230, 214, 52, 52, 0, 216, 
	216, 215, 57, 60, 70, 221, 216, 57, 
	57, 56, 216, 216, 215, 57, 60, 70, 
	221, 216, 57, 57, 56, 216, 216, 215, 
	57, 60, 70, 221, 216, 57, 57, 56, 
	216, 216, 215, 57, 60, 70, 221, 216, 
	57, 57, 56, 220, 220, 215, 59, 57, 
	60, 70, 221, 220, 57, 57, 57, 56, 
	220, 220, 215, 59, 57, 60, 70, 221, 
	220, 57, 57, 56, 221, 221, 215, 57, 
	60, 70, 221, 221, 57, 57, 56, 223, 
	223, 222, 87, 90, 100, 227, 223, 87, 
	87, 86, 223, 223, 222, 87, 90, 100, 
	227, 223, 87, 87, 86, 223, 223, 222, 
	87, 90, 100, 227, 223, 87, 87, 86, 
	226, 226, 222, 89, 87, 90, 100, 227, 
	226, 87, 87, 87, 86, 226, 226, 222, 
	89, 87, 90, 100, 227, 226, 87, 87, 
	86, 227, 227, 222, 87, 90, 100, 227, 
	227, 87, 87, 86, 229, 229, 54, 52, 
	75, 105, 230, 229, 52, 52, 52, 0, 
	229, 229, 54, 52, 75, 105, 230, 229, 
	52, 52, 0, 230, 230, 52, 75, 105, 
	230, 230, 52, 52, 0, 232, 232, 112, 
	110, 114, 121, 230, 232, 110, 110, 110, 
	0, 232, 232, 112, 52, 75, 105, 230, 
	232, 52, 52, 0, 0, 132, 133, 132, 
	169, 170, 169, 
}

var _trans_actions []byte = []byte{
	1, 1, 1, 26, 1, 1, 0, 41, 
	7, 7, 0, 7, 0, 0, 0, 0, 
	3, 0, 0, 0, 0, 3, 0, 0, 
	0, 0, 17, 20, 1, 3, 5, 0, 
	3, 0, 5, 1, 1, 1, 11, 0, 
	1, 1, 0, 29, 7, 5, 7, 0, 
	7, 0, 0, 0, 0, 3, 0, 5, 
	0, 0, 0, 3, 0, 5, 0, 0, 
	3, 5, 0, 29, 7, 5, 7, 0, 
	7, 0, 0, 0, 0, 3, 0, 5, 
	0, 0, 0, 3, 0, 5, 0, 0, 
	29, 7, 5, 7, 0, 0, 7, 0, 
	0, 0, 0, 29, 7, 5, 7, 0, 
	0, 7, 0, 0, 0, 0, 29, 7, 
	5, 7, 0, 0, 7, 0, 0, 0, 
	0, 29, 7, 5, 7, 0, 0, 7, 
	0, 0, 0, 0, 29, 7, 5, 7, 
	0, 0, 7, 0, 0, 0, 0, 29, 
	7, 5, 7, 1, 7, 0, 1, 1, 
	0, 3, 0, 5, 0, 1, 0, 1, 
	1, 0, 29, 7, 5, 7, 0, 0, 
	7, 0, 0, 0, 0, 29, 7, 5, 
	7, 0, 0, 7, 0, 0, 0, 0, 
	29, 7, 5, 7, 0, 0, 7, 0, 
	0, 0, 0, 29, 7, 5, 7, 0, 
	0, 7, 0, 0, 0, 0, 29, 7, 
	5, 7, 0, 0, 7, 0, 0, 0, 
	0, 3, 0, 5, 1, 1, 1, 11, 
	0, 1, 1, 0, 29, 7, 5, 7, 
	0, 0, 7, 0, 0, 0, 0, 29, 
	7, 5, 7, 0, 0, 7, 0, 0, 
	0, 0, 29, 7, 5, 7, 0, 0, 
	7, 0, 0, 0, 0, 29, 7, 5, 
	7, 0, 0, 7, 0, 0, 0, 0, 
	29, 7, 5, 7, 0, 0, 7, 0, 
	0, 0, 0, 29, 7, 5, 7, 1, 
	7, 0, 1, 1, 0, 3, 0, 5, 
	0, 1, 0, 1, 1, 0, 29, 7, 
	5, 7, 0, 0, 7, 0, 0, 0, 
	0, 29, 7, 5, 7, 0, 0, 7, 
	0, 0, 0, 0, 29, 7, 5, 7, 
	0, 0, 7, 0, 0, 0, 0, 29, 
	7, 5, 7, 0, 0, 7, 0, 0, 
	0, 0, 29, 7, 5, 7, 0, 0, 
	7, 0, 0, 0, 0, 41, 7, 7, 
	0, 0, 7, 0, 0, 0, 0, 41, 
	7, 7, 0, 0, 7, 0, 0, 0, 
	0, 41, 7, 7, 0, 0, 7, 0, 
	0, 0, 0, 41, 7, 7, 0, 0, 
	7, 0, 0, 0, 0, 41, 7, 7, 
	0, 0, 7, 0, 0, 0, 0, 41, 
	7, 7, 1, 7, 0, 1, 1, 0, 
	3, 0, 0, 1, 0, 1, 1, 0, 
	41, 7, 7, 0, 0, 7, 0, 0, 
	0, 0, 41, 7, 7, 0, 0, 7, 
	0, 0, 0, 0, 41, 7, 7, 0, 
	0, 7, 0, 0, 0, 0, 41, 7, 
	7, 0, 0, 7, 0, 0, 0, 0, 
	41, 7, 7, 0, 0, 7, 0, 0, 
	0, 0, 41, 7, 7, 0, 7, 0, 
	0, 0, 0, 3, 0, 0, 0, 0, 
	3, 0, 0, 0, 0, 17, 20, 1, 
	3, 5, 0, 41, 7, 5, 7, 0, 
	7, 0, 0, 0, 0, 3, 0, 5, 
	0, 0, 0, 3, 0, 5, 0, 0, 
	41, 7, 5, 7, 0, 0, 7, 0, 
	0, 0, 0, 41, 7, 5, 7, 0, 
	0, 7, 0, 0, 0, 0, 41, 7, 
	5, 7, 0, 0, 7, 0, 0, 0, 
	0, 41, 7, 5, 7, 0, 0, 7, 
	0, 0, 0, 0, 41, 7, 5, 7, 
	0, 0, 7, 0, 0, 0, 0, 41, 
	7, 5, 7, 1, 7, 0, 1, 1, 
	0, 3, 0, 5, 0, 1, 0, 1, 
	1, 0, 29, 7, 5, 7, 0, 7, 
	0, 0, 0, 0, 3, 0, 5, 0, 
	0, 0, 3, 0, 5, 0, 0, 41, 
	7, 5, 7, 0, 0, 7, 0, 0, 
	0, 0, 41, 7, 5, 7, 0, 0, 
	7, 0, 0, 0, 0, 41, 7, 5, 
	7, 0, 0, 7, 0, 0, 0, 0, 
	41, 7, 5, 7, 0, 0, 7, 0, 
	0, 0, 0, 41, 7, 5, 7, 0, 
	0, 7, 0, 0, 0, 0, 41, 7, 
	7, 0, 0, 7, 0, 0, 0, 0, 
	41, 7, 7, 0, 0, 7, 0, 0, 
	0, 0, 41, 7, 7, 0, 0, 7, 
	0, 0, 0, 0, 41, 7, 7, 0, 
	0, 7, 0, 0, 0, 0, 41, 7, 
	7, 0, 0, 7, 0, 0, 0, 0, 
	41, 7, 7, 1, 7, 0, 1, 1, 
	0, 3, 0, 0, 1, 0, 1, 1, 
	0, 41, 7, 7, 0, 7, 0, 0, 
	0, 0, 3, 0, 0, 0, 0, 3, 
	0, 0, 0, 0, 17, 20, 1, 3, 
	5, 0, 41, 7, 5, 7, 0, 7, 
	0, 0, 0, 0, 3, 0, 5, 0, 
	0, 0, 3, 0, 5, 0, 0, 41, 
	7, 5, 7, 0, 0, 7, 0, 0, 
	0, 0, 41, 7, 5, 7, 0, 0, 
	7, 0, 0, 0, 0, 41, 7, 5, 
	7, 0, 0, 7, 0, 0, 0, 0, 
	41, 7, 5, 7, 0, 0, 7, 0, 
	0, 0, 0, 41, 7, 5, 7, 0, 
	0, 7, 0, 0, 0, 0, 41, 7, 
	5, 7, 1, 7, 0, 1, 1, 0, 
	3, 0, 5, 0, 1, 0, 1, 1, 
	0, 41, 7, 5, 7, 0, 7, 0, 
	0, 0, 0, 3, 0, 5, 0, 0, 
	0, 3, 0, 5, 0, 0, 41, 7, 
	5, 7, 0, 0, 7, 0, 0, 0, 
	0, 41, 7, 5, 7, 0, 0, 7, 
	0, 0, 0, 0, 41, 7, 5, 7, 
	0, 0, 7, 0, 0, 0, 0, 41, 
	7, 5, 7, 0, 0, 7, 0, 0, 
	0, 0, 41, 7, 5, 7, 0, 0, 
	7, 0, 0, 0, 0, 41, 7, 7, 
	0, 0, 7, 0, 0, 0, 0, 41, 
	7, 7, 0, 0, 7, 0, 0, 0, 
	0, 41, 7, 7, 0, 0, 7, 0, 
	0, 0, 0, 41, 7, 7, 0, 0, 
	7, 0, 0, 0, 0, 41, 7, 7, 
	0, 0, 7, 0, 0, 0, 0, 41, 
	7, 7, 0, 7, 0, 0, 0, 0, 
	3, 0, 0, 0, 0, 3, 0, 0, 
	0, 0, 17, 20, 1, 41, 7, 7, 
	0, 0, 7, 0, 0, 0, 0, 41, 
	7, 7, 0, 0, 7, 0, 0, 0, 
	0, 41, 7, 7, 0, 0, 7, 0, 
	0, 0, 0, 41, 7, 7, 0, 0, 
	7, 0, 0, 0, 0, 41, 7, 7, 
	0, 0, 7, 0, 0, 0, 0, 41, 
	7, 7, 1, 7, 0, 1, 1, 0, 
	3, 0, 0, 1, 0, 1, 1, 0, 
	41, 7, 7, 0, 0, 7, 0, 0, 
	0, 0, 41, 7, 7, 0, 0, 7, 
	0, 0, 0, 0, 41, 7, 7, 0, 
	0, 7, 0, 0, 0, 0, 41, 7, 
	7, 0, 0, 7, 0, 0, 0, 0, 
	41, 7, 7, 0, 0, 7, 0, 0, 
	0, 0, 3, 0, 1, 1, 1, 11, 
	13, 0, 1, 1, 0, 3, 0, 13, 
	0, 0, 41, 7, 7, 0, 7, 0, 
	0, 0, 0, 3, 0, 0, 0, 0, 
	3, 0, 0, 0, 0, 17, 20, 1, 
	3, 5, 0, 32, 9, 5, 44, 44, 
	44, 47, 50, 9, 44, 44, 0, 3, 
	0, 5, 13, 0, 0, 41, 7, 5, 
	7, 0, 7, 0, 0, 0, 0, 3, 
	0, 5, 0, 0, 0, 3, 0, 5, 
	0, 0, 56, 23, 20, 23, 23, 23, 
	60, 64, 23, 23, 23, 1, 41, 7, 
	5, 7, 0, 0, 7, 0, 0, 0, 
	0, 41, 7, 5, 7, 0, 0, 7, 
	0, 0, 0, 0, 41, 7, 5, 7, 
	0, 0, 7, 0, 0, 0, 0, 41, 
	7, 5, 7, 0, 0, 7, 0, 0, 
	0, 0, 41, 7, 5, 7, 0, 0, 
	7, 0, 0, 0, 0, 41, 7, 5, 
	7, 1, 7, 0, 1, 1, 0, 3, 
	0, 5, 0, 1, 0, 1, 1, 0, 
	29, 7, 5, 7, 0, 7, 0, 0, 
	0, 0, 3, 0, 5, 0, 0, 0, 
	3, 0, 5, 0, 0, 68, 44, 38, 
	44, 44, 44, 72, 76, 44, 44, 44, 
	1, 41, 7, 5, 7, 0, 0, 7, 
	0, 0, 0, 0, 41, 7, 5, 7, 
	0, 0, 7, 0, 0, 0, 0, 41, 
	7, 5, 7, 0, 0, 7, 0, 0, 
	0, 0, 41, 7, 5, 7, 0, 0, 
	7, 0, 0, 0, 0, 41, 7, 5, 
	7, 0, 0, 7, 0, 0, 0, 0, 
	41, 7, 5, 7, 1, 1, 1, 11, 
	13, 7, 0, 1, 1, 0, 3, 0, 
	5, 0, 1, 1, 1, 11, 13, 0, 
	1, 1, 0, 35, 11, 5, 1, 1, 
	1, 11, 13, 11, 1, 1, 0, 41, 
	7, 7, 0, 0, 7, 0, 0, 0, 
	0, 41, 7, 7, 0, 0, 7, 0, 
	0, 0, 0, 41, 7, 7, 0, 0, 
	7, 0, 0, 0, 0, 41, 7, 7, 
	0, 0, 7, 0, 0, 0, 0, 41, 
	7, 7, 0, 0, 7, 0, 0, 0, 
	0, 41, 7, 7, 1, 7, 0, 1, 
	1, 0, 3, 0, 0, 1, 0, 1, 
	1, 0, 41, 7, 7, 0, 7, 0, 
	0, 0, 0, 3, 0, 0, 0, 0, 
	3, 0, 0, 0, 0, 17, 20, 1, 
	3, 5, 0, 32, 9, 5, 44, 44, 
	44, 47, 50, 9, 44, 44, 0, 3, 
	0, 5, 13, 0, 0, 41, 7, 5, 
	7, 0, 7, 0, 0, 0, 0, 3, 
	0, 5, 0, 0, 0, 3, 0, 5, 
	0, 0, 41, 7, 5, 7, 0, 0, 
	7, 0, 0, 0, 0, 41, 7, 5, 
	7, 0, 0, 7, 0, 0, 0, 0, 
	41, 7, 5, 7, 0, 0, 7, 0, 
	0, 0, 0, 41, 7, 5, 7, 0, 
	0, 7, 0, 0, 0, 0, 41, 7, 
	5, 7, 0, 0, 7, 0, 0, 0, 
	0, 41, 7, 5, 7, 1, 7, 0, 
	1, 1, 0, 3, 0, 5, 0, 1, 
	0, 1, 1, 0, 41, 7, 5, 7, 
	0, 7, 0, 0, 0, 0, 3, 0, 
	5, 0, 0, 0, 3, 0, 5, 0, 
	0, 56, 23, 20, 23, 23, 23, 60, 
	64, 23, 23, 23, 1, 41, 7, 5, 
	7, 0, 0, 7, 0, 0, 0, 0, 
	41, 7, 5, 7, 0, 0, 7, 0, 
	0, 0, 0, 41, 7, 5, 7, 0, 
	0, 7, 0, 0, 0, 0, 41, 7, 
	5, 7, 0, 0, 7, 0, 0, 0, 
	0, 41, 7, 5, 7, 0, 0, 7, 
	0, 0, 0, 0, 41, 7, 5, 7, 
	1, 1, 1, 11, 13, 7, 0, 1, 
	1, 0, 3, 0, 5, 0, 1, 1, 
	1, 11, 13, 0, 1, 1, 0, 35, 
	11, 5, 1, 1, 1, 11, 13, 11, 
	1, 1, 0, 41, 7, 7, 0, 0, 
	7, 0, 0, 0, 0, 41, 7, 7, 
	0, 0, 7, 0, 0, 0, 0, 41, 
	7, 7, 0, 0, 7, 0, 0, 0, 
	0, 41, 7, 7, 0, 0, 7, 0, 
	0, 0, 0, 41, 7, 7, 0, 0, 
	7, 0, 0, 0, 0, 41, 7, 7, 
	1, 1, 1, 11, 13, 7, 0, 1, 
	1, 0, 3, 0, 0, 1, 1, 1, 
	11, 13, 0, 1, 1, 0, 35, 11, 
	1, 1, 1, 11, 13, 11, 1, 1, 
	0, 32, 9, 5, 44, 44, 44, 47, 
	9, 44, 44, 0, 68, 44, 38, 44, 
	44, 44, 72, 44, 44, 44, 1, 32, 
	9, 5, 44, 44, 44, 47, 9, 44, 
	44, 0, 3, 0, 5, 1, 1, 1, 
	11, 0, 1, 1, 0, 68, 44, 38, 
	44, 44, 44, 72, 44, 44, 44, 1, 
	29, 7, 5, 7, 1, 1, 1, 11, 
	7, 0, 1, 1, 0, 3, 0, 5, 
	0, 1, 1, 1, 11, 0, 1, 1, 
	0, 35, 11, 5, 1, 1, 1, 11, 
	11, 1, 1, 0, 29, 7, 5, 7, 
	1, 1, 1, 11, 7, 0, 1, 1, 
	0, 3, 0, 5, 0, 1, 1, 1, 
	11, 0, 1, 1, 0, 35, 11, 5, 
	1, 1, 1, 11, 11, 1, 1, 0, 
	41, 7, 7, 1, 1, 1, 11, 7, 
	0, 1, 1, 0, 3, 0, 0, 1, 
	1, 1, 11, 0, 1, 1, 0, 32, 
	9, 5, 44, 44, 44, 47, 9, 44, 
	44, 0, 3, 0, 5, 1, 1, 1, 
	11, 0, 1, 1, 0, 56, 23, 20, 
	23, 23, 23, 60, 23, 23, 23, 1, 
	68, 44, 38, 44, 44, 44, 72, 44, 
	44, 44, 1, 41, 7, 5, 7, 1, 
	1, 1, 11, 7, 0, 1, 1, 0, 
	3, 0, 5, 0, 1, 1, 1, 11, 
	0, 1, 1, 0, 35, 11, 5, 1, 
	1, 1, 11, 11, 1, 1, 0, 32, 
	9, 5, 44, 44, 44, 47, 9, 44, 
	44, 0, 3, 0, 5, 1, 1, 1, 
	11, 0, 1, 1, 0, 56, 23, 20, 
	23, 23, 23, 60, 23, 23, 23, 1, 
	41, 7, 5, 7, 1, 1, 1, 11, 
	7, 0, 1, 1, 0, 3, 0, 5, 
	0, 1, 1, 1, 11, 0, 1, 1, 
	0, 35, 11, 5, 1, 1, 1, 11, 
	11, 1, 1, 0, 41, 7, 7, 1, 
	1, 1, 11, 7, 0, 1, 1, 0, 
	3, 0, 0, 1, 1, 1, 11, 0, 
	1, 1, 0, 35, 11, 1, 1, 1, 
	11, 11, 1, 1, 0, 41, 7, 7, 
	1, 1, 1, 11, 7, 0, 1, 1, 
	0, 3, 0, 0, 1, 1, 1, 11, 
	0, 1, 1, 0, 0, 3, 5, 0, 
	3, 5, 0, 
}

var _eof_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 53, 53, 53, 15, 53, 15, 
	15, 15, 15, 15, 15, 15, 15, 53, 
	15, 53, 53, 15, 15, 15, 53, 15, 
	53, 15, 15, 15, 15, 15, 15, 15, 
	15, 0, 0, 0, 
}

const start int = 1
const first_final int = 202

const en_closingbrace int = 126
const en_main int = 1


//line /tmp/flux.rl:176

type machine struct {
	data         []byte
    stack        []int
    top          int
	cs           int
	p, pe, eof   int
	pb           int
    curline      int
	sol 		 int // Position of the start of line
	err          error
	root 	     *ast.Program
	statements   []ast.Statement 	 
	identifier   *ast.Identifier
	expression   ast.Expression
}

func NewMachine() *machine {
	m := &machine{}

	
//line /tmp/flux.rl:197
	
//line /tmp/flux.rl:198
	
//line /tmp/flux.rl:199
	
//line /tmp/flux.rl:200
	
//line /tmp/flux.rl:201
    
//line /tmp/flux.rl:202
    
//line /tmp/flux.rl:203

	return m
}

func (m *machine) text() []byte {
	return m.data[m.pb:m.p]
}

func (m *machine) col() int {
	return m.p - m.sol
}

func (m *machine) Parse(input []byte) *ast.Program {
	m.data = input
    m.curline = 1
	m.sol = 0
	m.p = 0
	m.pb = 0
    m.top = 0
    m.stack = nil
	m.pe = len(input)
	m.eof = len(input)
	m.err = nil

    
//line /tmp/flux.go:1162
	{
	 m.cs = start
	( m.top) = 0
	}

//line /tmp/flux.rl:228
    
//line /tmp/flux.go:1170
	{
	var _klen int
	var _trans int
	var _acts int
	var _nacts uint
	var _keys int
	if ( m.p) == ( m.pe) {
		goto _test_eof
	}
	if  m.cs == 0 {
		goto _out
	}
_resume:
	_keys = int(_key_offsets[ m.cs])
	_trans = int(_index_offsets[ m.cs])

	_klen = int(_single_lengths[ m.cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + _klen - 1)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + ((_upper - _lower) >> 1)
			switch {
			case ( m.data)[( m.p)] < _trans_keys[_mid]:
				_upper = _mid - 1
			case ( m.data)[( m.p)] > _trans_keys[_mid]:
				_lower = _mid + 1
			default:
				_trans += int(_mid - int(_keys))
				goto _match
			}
		}
		_keys += _klen
		_trans += _klen
	}

	_klen = int(_range_lengths[ m.cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + (_klen << 1) - 2)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + (((_upper - _lower) >> 1) & ^1)
			switch {
			case ( m.data)[( m.p)] < _trans_keys[_mid]:
				_upper = _mid - 2
			case ( m.data)[( m.p)] > _trans_keys[_mid + 1]:
				_lower = _mid + 2
			default:
				_trans += int((_mid - int(_keys)) >> 1)
				goto _match
			}
		}
		_trans += _klen
	}

_match:
	 m.cs = int(_trans_targs[_trans])

	if _trans_actions[_trans] == 0 {
		goto _again
	}

	_acts = int(_trans_actions[_trans])
	_nacts = uint(_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _actions[_acts-1] {
		case 0:
//line /tmp/flux.rl:14

	fmt.Println("mark", m.pb, m.p, m.stack)
	m.pb = m.p

		case 1:
//line /tmp/flux.rl:19

	fmt.Println("INSIDEN")
	m.sol = m.p
    m.curline++;

		case 2:
//line /tmp/flux.rl:38

	m.expression = &ast.StringLiteral{
		Value: string(m.text()),
	}

		case 3:
//line /tmp/flux.rl:90

	m.identifier = &ast.Identifier{
		Name: string(m.text()),
		// BaseNode: base(m.text(), m.curline, m.col()),
	}

		case 4:
//line /tmp/flux.rl:97

	m.statements = append(m.statements, &ast.VariableDeclaration{
		Declarations: []*ast.VariableDeclarator{
			&ast.VariableDeclarator{
				ID:   m.identifier,
				Init: m.expression,
			},
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil
	m.expression = nil

		case 5:
//line /tmp/flux.rl:125
 { 
	m.stack = append(m.stack, 0)
( m.stack)[( m.top)] =  m.cs; ( m.top)++;  m.cs = 126; goto _again
 } 
		case 6:
//line /tmp/flux.rl:127
 ( m.top)--;  m.cs = ( m.stack)[( m.top)]
{ 
	m.stack = m.stack[:len(m.stack) - 1]
 }
goto _again
 
//line /tmp/flux.go:1305
		}
	}

_again:
	if  m.cs == 0 {
		goto _out
	}
	( m.p)++
	if ( m.p) != ( m.pe) {
		goto _resume
	}
	_test_eof: {}
	if ( m.p) == ( m.eof) {
		__acts := _eof_actions[ m.cs]
		__nacts := uint(_actions[__acts]); __acts++
		for ; __nacts > 0; __nacts-- {
			__acts++
			switch _actions[__acts-1] {
			case 4:
//line /tmp/flux.rl:97

	m.statements = append(m.statements, &ast.VariableDeclaration{
		Declarations: []*ast.VariableDeclarator{
			&ast.VariableDeclarator{
				ID:   m.identifier,
				Init: m.expression,
			},
		},
		// BaseNode: base(m.text(), m.curline, m.col()),
	})
	m.identifier = nil
	m.expression = nil

			case 7:
//line /tmp/flux.rl:137

	fmt.Println("ex_program")

	m.root = &ast.Program{
		Body:     m.statements,
		// BaseNode: base(m.text(), m.curline, m.col()),
	}

	// m.children = nil

	fmt.Println(m.p)
	spew.Dump(m.root)

//line /tmp/flux.go:1354
			}
		}
	}

	_out: {}
	}

//line /tmp/flux.rl:229

	if m.cs < first_final  {
		return nil
	}

	return m.root
}