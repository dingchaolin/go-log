/**
 * Created by chaolinding on 2018/3/22.
 */

package test

import (
	"runtime/debug"
	"errors"
	"fmt"
	"os"
)

func Test(){
	err := errors.New("我是错误")
	if err != nil {
		fmt.Fprint( os.Stderr, err.Error() )
		debug.PrintStack()
	}

}