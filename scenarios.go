package main

import (
    "go.uber.org/zap"

    "github.com/yandex/pandora/core"
    "github.com/yandex/pandora/core/aggregate/netsample"
)

func (l *Gun) scenario1(a core.Ammo) {
    sample := netsample.Acquire("REQUEST")
    // Do work here.
    l.Log.Info("Example Gun message", zap.String("message", a.(*Ammo).Message))
    sample.SetProtoCode(200)
    l.aggregator.Report(sample)
}