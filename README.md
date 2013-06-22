# beelog

A simple log for Debug Trace Warn Error Critacal, extract code from astaxie/beego.


## 日志处理

beelog默认有一个初始化的BeeLogger对象输出内容到stdout中，你可以通过如下的方式设置自己的输出：

	beelog.SetLogger(*log.Logger)

只要你的输出符合`*log.Logger`就可以，例如输出到文件：

	fd,err := os.OpenFile("/var/log/beeapp/beeapp.log", os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		beelog.Critical("openfile beeapp.log:", err)
		return
	}
	lg := log.New(fd, "", log.Ldate|log.Ltime)
	beelog.SetLogger(lg)


### 不同级别的log日志函数

* Trace(v ...interface{})
* Debug(v ...interface{})
* Info(v ...interface{})
* Warn(v ...interface{})
* Error(v ...interface{})
* Critical(v ...interface{})

你可以通过下面的方式设置不同的日志分级：

	beelog.SetLevel(beelog.LevelError)

当你代码中有很多日志输出之后，如果想上线，但是你不想输出Trace、Debug、Info等信息，那么你可以设置如下：

	beelog.SetLevel(beelog.LevelWarning)

这样的话就不会输出小于这个level的日志，日志的排序如下：

LevelTrace、LevelDebug、LevelInfo、LevelWarning、LevelError、LevelCritical

用户可以根据不同的级别输出不同的错误信息，如下例子所示：


### Examples of log messages

- Trace

	* "Entered parse function validation block"
	* "Validation: entered second 'if'"
	* "Dictionary 'Dict' is empty. Using default value"

- Debug

	* "Web page requested: http://somesite.com Params='...'"
	* "Response generated. Response size: 10000. Sending."
	* "New file received. Type:PNG Size:20000"

- Info

	* "Web server restarted"
	* "Hourly statistics: Requested pages: 12345 Errors: 123 ..."
	* "Service paused. Waiting for 'resume' call"

- Warn

	* "Cache corrupted for file='test.file'. Reading from back-end"
	* "Database 192.168.0.7/DB not responding. Using backup 192.168.0.8/DB"
	* "No response from statistics server. Statistics not sent"

- Error

	* "Internal error. Cannot process request #12345 Error:...."
	* "Cannot perform login: credentials DB not responding"

- Critical

	* "Critical panic received: .... Shutting down"
	* "Fatal error: ... App is shutting down to prevent data corruption or loss"


### Example

	func internalCalculationFunc(x, y int) (result int, err error) {
		beelog.Debug("calculating z. x:", x, " y:", y)
		z := y
		switch {
		case x == 3:
			beelog.Trace("x == 3")
			panic("Failure.")
		case y == 1:
			beelog.Trace("y == 1")
			return 0, errors.New("Error!")
		case y == 2:
			beelog.Trace("y == 2")
			z = x
		default:
			beelog.Trace("default")
			z += x
		}
		retVal := z - 3
		beelog.Debug("Returning ", retVal)
		
		return retVal, nil
	}
	
	func processInput(input inputData) {
		defer func() {
			if r := recover(); r != nil {
				beelog.Error("Unexpected error occurred: ", r)
				outputs <- outputData{result: 0, error: true}
			}
		}()
		beelog.Info("Received input signal. x:", input.x, " y:", input.y)
		
		res, err := internalCalculationFunc(input.x, input.y)
		if err != nil {
			beelog.Warn("Error in calculation:", err.Error())
		}
		
		beelog.Info("Returning result: ", res, " error: ", err)
		outputs <- outputData{result: res, error: err != nil}
	}
	
	func main() {
		inputs = make(chan inputData)
		outputs = make(chan outputData)
		criticalChan = make(chan int)
		beelog.Info("App started.")
		
		go consumeResults(outputs)
		beelog.Info("Started receiving results.")
		
		go generateInputs(inputs)
		beelog.Info("Started sending signals.")
		
		for {
			select {
			case input := <-inputs:
				processInput(input)
			case <-criticalChan:
				beelog.Critical("Caught value from criticalChan: Go shut down.")
				panic("Shut down due to critical fault.")
			}
		}
	}
