package runenv

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRunEnvKey(t *testing.T) {
	Convey("[TestGetRunEnvKey] GetRunEnv", t, func() {
		So(GetRunEnvKey(), ShouldEqual, runEnvKey)
	})

	Convey("[TestGetRunEnvKey] SetRunEnvKey", t, func() {
		newKey := "RUN_ENV"
		So(GetRunEnvKey(), ShouldEqual, runEnvKey)
		So(SetRunEnvKey(newKey), ShouldBeNil)
		So(GetRunEnvKey(), ShouldEqual, newKey)

		So(SetRunEnvKey(""), ShouldNotBeNil)
		So(GetRunEnvKey(), ShouldEqual, newKey)
	})
}

func TestRunEnv(t *testing.T) {
	Convey("[TestRunEnv]", t, func() {

		// default
		_ = os.Setenv(GetRunEnvKey(), "")
		So(GetRunEnv(), ShouldEqual, Dev)

		allRunEnv := []REnv{Dev, Test, Gray, Prod}

		for _, rEnv := range allRunEnv {

			_ = os.Setenv(GetRunEnvKey(), rEnv)
			So(GetRunEnv(), ShouldEqual, rEnv)

			var isDev, isTest, isGray, isProd bool

			switch rEnv {
			case Dev:
				isDev = true
			case Test:
				isTest = true
			case Gray:
				isGray = true
			case Prod:
				isProd = true
			}

			So(IsDev(), ShouldEqual, isDev)
			So(IsTest(), ShouldEqual, isTest)
			So(IsGray(), ShouldEqual, isGray)
			So(IsProd(), ShouldEqual, isProd)

			So(Is(Dev), ShouldEqual, isDev)
			So(Not(Dev), ShouldNotEqual, isDev)
			So(Is(Test), ShouldEqual, isTest)
			So(Not(Test), ShouldNotEqual, isTest)
			So(Is(Gray), ShouldEqual, isGray)
			So(Not(Gray), ShouldNotEqual, isGray)
			So(Is(Prod), ShouldEqual, isProd)
			So(Not(Prod), ShouldNotEqual, isProd)
		}
	})
}
