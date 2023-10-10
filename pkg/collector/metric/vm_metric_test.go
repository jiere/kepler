package metric

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("VMMetric", func() {

	It("Test ResetDeltaValues", func() {
		p := NewVMMetrics(0, "name")
		p.ResetDeltaValues()
		Expect(p.CPUTime.Delta).To(Equal(uint64(0)))
	})
})