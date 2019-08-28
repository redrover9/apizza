package dawg

import (
	"testing"
)

// Move this to an items_test.go file
func TestItems(t *testing.T) {
	store := testingStore()
	menu, err := store.Menu()
	if err != nil {
		t.Error(err)
	}

	testcases := []struct {
		product, variant, top, cover, coverErr string
		isSubset, wanterr                      bool
	}{
		// {product: "F_PARMT", variant: "B8PCPT", top: "K", isSubset: true, wanterr: false},
		{
			product:  "S_MX",
			variant:  "14TMEATZA",
			top:      "B",
			isSubset: true,
			wanterr:  false,
			cover:    "2",
			coverErr: "1.7",
		},
		{
			product:  "S_PISPF",
			variant:  "P10IRESPF",
			top:      "B",
			isSubset: true,
			wanterr:  false,
			cover:    "2",
			coverErr: "-1.7",
		},
		{
			product:  "S_BONELESS",
			variant:  "W08PBNLW",
			top:      "",
			isSubset: true,
			wanterr:  false,
			cover:    "2",
			coverErr: "-1.7",
		},
	}

	for _, tc := range testcases {
		p, err := menu.GetProduct(tc.product)
		if tc.wanterr && err == nil {
			t.Error("expected error")
		} else if err != nil {
			t.Error(err)
		}
		v, err := menu.GetVariant(tc.variant)
		if tc.wanterr && err == nil {
			t.Error("expected error")
		} else if err != nil {
			t.Error(err)
		}

		if tc.isSubset {
			for _, variant := range p.Variants {
				if variant == tc.variant {
					goto foundVariant
				}
			}
			t.Errorf("%s should be a variant of %s", tc.variant, tc.product)
		foundVariant:
		}
		if err = p.AddTopping(tc.top, ToppingLeft, tc.cover); err != nil {
			t.Error(err)
		}
		if err = v.AddTopping(tc.top, ToppingFull, tc.cover); err != nil {
			t.Error(err)
		}
		if err = v.AddTopping(tc.top, "1/1", tc.coverErr); err == nil {
			t.Error("expected error")
		}
		if len(v.opts) < 1 {
			t.Error("should have options in the struct")
		}
	}
}

func TestOPFromItem(t *testing.T) {
	m := testingMenu()
	v, err := m.GetVariant("W08PBNLW")
	if err != nil {
		t.Error(err)
	}
	p, err := m.GetProduct("S_BONELESS")
	if err != nil {
		t.Error(err)
	}

	opv := OrderProductFromItem(v)
	opp := OrderProductFromItem(p)

	opvOpts := opv.Options()
	oppOpts := opp.Options()

	for k := range opvOpts {
		if _, ok := oppOpts[k]; !ok {
			t.Errorf("order product should have %s", k)
		}
	}
	for k := range v.Options() {
		if _, ok := opvOpts[k]; !ok {
			t.Error("options should be the same")
		}
	}
	for k := range p.Options() {
		if _, ok := oppOpts[k]; !ok {
			t.Error("options should be the same")
		}
	}

	if opv.Category() != opp.Category() {
		t.Error("the variant and it's parent should have the same product type")
	}
}

func TestFindItem(t *testing.T) {
	m := testingMenu()
	itm := m.FindItem("W08PBNLW")
	if itm == nil {
		t.Error("item is nil")
	}
	itm = m.FindItem("S_BONELESS")
	if itm == nil {
		t.Error("item is nil")
	}
	itm = m.FindItem("F_PARMT")
	if itm == nil {
		t.Error("item is nil")
	}
	itm = m.FindItem("P_14SCREEN")
	if itm == nil {
		t.Error("item is nil")
	}

	itm = m.FindItem("badCode")
	if itm != nil {
		t.Error("item should be nil")
	}
}
