package chinese

import "testing"

func TestIsChinese(t *testing.T) {
	if IsChinese("年") == false {
		t.Errorf("failed")
	}

	if IsChinese("月") == false {
		t.Errorf("failed")
	}

	if IsChinese("日") == false {
		t.Errorf("failed")
	}

	if IsChinese("时") == false {
		t.Errorf("failed")
	}

	if IsChinese("分") == false {
		t.Errorf("failed")
	}

	if IsChinese("秒") == false {
		t.Errorf("failed")
	}

	if IsChinese("2006") == true {
		t.Errorf("failed")
	}
}

func TestIsWithoutTime(t *testing.T) {
	if IsWithoutTime("") == false {
		t.Errorf("failed")
	}

	if IsWithoutTime("2006年01月02日") == false {
		t.Errorf("failed")
	}

	if IsWithoutTime("2006年01月02日 15时04分05秒") == true {
		t.Errorf("failed")
	}
}

func TestIsAbbrYear(t *testing.T) {
	if IsAbbrYear("2006年01月02日 15时04分05秒") == true {
		t.Errorf("failed")
	}

	if IsAbbrYear("06年01月02日 15时04分05秒") == false {
		t.Errorf("failed")
	}
}

func TestIsWithoutSecond(t *testing.T) {
	if IsWithoutSecond("2006年01月02日 15时04分05秒") == true {
		t.Errorf("failed")
	}

	if IsWithoutSecond("2006年01月02日 15时04分") == false {
		t.Errorf("failed")
	}
}
