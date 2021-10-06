package repo

import (
	"bytes"
	"io"
	"testing"

	geojson "github.com/paulmach/go.geojson"
	"github.com/pkg/errors"
)

// TestRepoSave tests save functionality
func TestRepoSave(t *testing.T) {
	r, err := NewSaveRepo()
	if err != nil {
		t.Fatal(errors.Wrap(err, "Error creating save repo"))
	}

	reader, writer := io.Pipe()
	var b bytes.Buffer
	go func() {
		writer.Write([]byte(`{"type":"Feature","geometry":{"type":"Polygon","coordinates":[[[143.504228803,43.844554794],[143.504209418,43.844068864],[143.504422486,43.844058515],[143.504663297,43.844013026],[143.50489573,43.843947557],[143.505126201,43.843877518],[143.505364127,43.843824949],[143.50558229,43.843725522],[143.505747581,43.843500842],[143.509296966,43.84271713],[143.509345793,43.842809482],[143.509393148,43.842908691],[143.509450871,43.843033708],[143.509485048,43.843121271],[143.509515538,43.843205004],[143.509542918,43.84328283],[143.509576339,43.843393534],[143.506032114,43.844175937],[143.505764871,43.844158982],[143.505553774,43.844275164],[143.505326575,43.844352998],[143.505074636,43.844371968],[143.504835528,43.844421669],[143.504605191,43.844492245],[143.504228803,43.844554794]]]},"properties":{"A33_001":"2","A33_002":"1","A33_003":"01","A33_004":"Ⅱ-75-1060","A33_005":"堀江の沢川","A33_006":"紋別郡遠軽町清里","A33_007":"2017/3/10","A33_008":"0"}}`))
		writer.Write([]byte("\n"))
		writer.Write([]byte(`{"type":"Feature","geometry":{"type":"Polygon","coordinates":[[[143.508490942,43.831127714],[143.508740604,43.83080705],[143.508990117,43.830873878],[143.509193887,43.830977125],[143.509407365,43.831019597],[143.509653747,43.831061346],[143.509933695,43.831048089],[143.510232834,43.831086464],[143.510832063,43.831274005],[143.510779359,43.831416352],[143.510699836,43.831574117],[143.510561203,43.831768987],[143.510421095,43.831918509],[143.509844032,43.831737382],[143.50960809,43.831581026],[143.509389681,43.831493424],[143.509177261,43.831396124],[143.508870036,43.831297618],[143.508645513,43.831214942],[143.508490942,43.831127714]]]},"properties":{"A33_001":"2","A33_002":"1","A33_003":"01","A33_004":"Ⅱ-75-1050","A33_005":"拓北沢川","A33_006":"紋別郡遠軽町清里","A33_007":"2017/3/10","A33_008":"0"}}`))
		writer.Write([]byte("\n"))
		writer.Close()
	}()
	r.Save(reader, &b)

	str := b.String()
	fc, err := geojson.UnmarshalFeatureCollection([]byte(str))
	if err != nil {
		t.Error(errors.Wrap(err, "Error marshalling feature collection"))
	}
	if fc.Type != "FeatureCollection" {
		t.Error(errors.Wrapf(err, "Expected %s but got %s", "FeatureCollection", fc.Type))
	}
	if len(fc.Features) != 2 {
		t.Error(errors.Wrapf(err, "Expected %d items but got %d", 2, len(fc.Features)))
	}
}

// TestRepoSavePath tests save functionality
func TestRepoSavePath(t *testing.T) {
	r, err := NewSaveRepo()
	if err != nil {
		t.Fatal(errors.Wrap(err, "Error creating save repo"))
	}

	reader, writer := io.Pipe()
	var b bytes.Buffer
	go func() {
		writer.Write([]byte(`{"type":"Feature","geometry":{"type":"Polygon","coordinates":[[[143.504228803,43.844554794],[143.504209418,43.844068864],[143.504422486,43.844058515],[143.504663297,43.844013026],[143.50489573,43.843947557],[143.505126201,43.843877518],[143.505364127,43.843824949],[143.50558229,43.843725522],[143.505747581,43.843500842],[143.509296966,43.84271713],[143.509345793,43.842809482],[143.509393148,43.842908691],[143.509450871,43.843033708],[143.509485048,43.843121271],[143.509515538,43.843205004],[143.509542918,43.84328283],[143.509576339,43.843393534],[143.506032114,43.844175937],[143.505764871,43.844158982],[143.505553774,43.844275164],[143.505326575,43.844352998],[143.505074636,43.844371968],[143.504835528,43.844421669],[143.504605191,43.844492245],[143.504228803,43.844554794]]]},"properties":{"A33_001":"2","A33_002":"1","A33_003":"01","A33_004":"Ⅱ-75-1060","A33_005":"堀江の沢川","A33_006":"紋別郡遠軽町清里","A33_007":"2017/3/10","A33_008":"0"}}`))
		writer.Write([]byte("\n"))
		writer.Write([]byte(`{"type":"Feature","geometry":{"type":"Polygon","coordinates":[[[143.508490942,43.831127714],[143.508740604,43.83080705],[143.508990117,43.830873878],[143.509193887,43.830977125],[143.509407365,43.831019597],[143.509653747,43.831061346],[143.509933695,43.831048089],[143.510232834,43.831086464],[143.510832063,43.831274005],[143.510779359,43.831416352],[143.510699836,43.831574117],[143.510561203,43.831768987],[143.510421095,43.831918509],[143.509844032,43.831737382],[143.50960809,43.831581026],[143.509389681,43.831493424],[143.509177261,43.831396124],[143.508870036,43.831297618],[143.508645513,43.831214942],[143.508490942,43.831127714]]]},"properties":{"A33_001":"2","A33_002":"1","A33_003":"01","A33_004":"Ⅱ-75-1050","A33_005":"拓北沢川","A33_006":"紋別郡遠軽町清里","A33_007":"2017/3/10","A33_008":"0"}}`))
		writer.Write([]byte("\n"))
		writer.Close()
	}()
	r.Save(reader, &b)

	str := b.String()
	fc, err := geojson.UnmarshalFeatureCollection([]byte(str))
	if err != nil {
		t.Error(errors.Wrap(err, "Error marshalling feature collection"))
	}
	if fc.Type != "FeatureCollection" {
		t.Error(errors.Wrapf(err, "Expected %s but got %s", "FeatureCollection", fc.Type))
	}
	if len(fc.Features) != 2 {
		t.Error(errors.Wrapf(err, "Expected %d items but got %d", 2, len(fc.Features)))
	}
}
