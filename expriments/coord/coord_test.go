package coord

import (
	"fmt"
	"testing"

	"github.com/qichengzx/coordtransform"
)

var src = [][2]float64{
	[2]float64{116.404123, 39.915321},
	[2]float64{114.21892734521, 29.575429778924},
	[2]float64{116.486929, 39.962983},
}

var dst = [][2]float64{
	[2]float64{116.4166272438, 39.9226995522},
}

func TestCoord(t *testing.T) {
	fmt.Println(coordTransform.WGS84toBD09(114.21892734521, 29.575429778924))
	for _, p := range src {
		lng, lat := coordTransform.WGS84toBD09(p[0], p[1])
		fmt.Printf("%.10f,%.10f\n", lng, lat)
	}
}

/*
116.4167504544,39.9230193316
116.4167559420,39.9230235927

114.2307469711,29.5790799570


116.4994774851,39.9703014026

*/
