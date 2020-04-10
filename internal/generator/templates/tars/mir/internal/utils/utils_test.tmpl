package utils

import "testing"

func TestObjFrom(t *testing.T) {
	for i, e := range []struct {
		name   string
		points []string
		expect string
	}{
		{
			name: "kry.calculator.CalculatorObj",
			points: []string{
				"tcp -h 127.0.0.1 -p 10028 -t 60000",
			},
			expect: "kry.calculator.CalculatorObj@tcp -h 127.0.0.1 -p 10028 -t 60000",
		},
		{
			name: "a.b.c",
			points: []string{
				"tcp -h 127.0.0.1 -p 10028 -t 60000",
				"tcp -h localhost -p 10029",
				"tcp -h 172.0.0.1 -p 10030",
			},
			expect: "a.b.c@tcp -h 127.0.0.1 -p 10028 -t 60000:tcp -h localhost -p 10029:tcp -h 172.0.0.1 -p 10030",
		},
	} {
		if res := ObjFrom(e.name, e.points...); res != e.expect {
			t.Errorf("[%d] want %s but get %s", i, e.expect, res)
		}
	}
}
