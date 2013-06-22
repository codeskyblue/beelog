# beelog

A simple log for Debug Trace Warn Error Critacal, extract code from astaxie/beego.


## ��־����

beelogĬ����һ����ʼ����BeeLogger����������ݵ�stdout�У������ͨ�����µķ�ʽ�����Լ��������

	beelog.SetLogger(*log.Logger)

ֻҪ����������`*log.Logger`�Ϳ��ԣ�����������ļ���

	fd,err := os.OpenFile("/var/log/beeapp/beeapp.log", os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		beelog.Critical("openfile beeapp.log:", err)
		return
	}
	lg := log.New(fd, "", log.Ldate|log.Ltime)
	beelog.SetLogger(lg)


### ��ͬ�����log��־����

* Trace(v ...interface{})
* Debug(v ...interface{})
* Info(v ...interface{})
* Warn(v ...interface{})
* Error(v ...interface{})
* Critical(v ...interface{})

�����ͨ������ķ�ʽ���ò�ͬ����־�ּ���

	beelog.SetLevel(beelog.LevelError)

����������кܶ���־���֮����������ߣ������㲻�����Trace��Debug��Info����Ϣ����ô������������£�

	beelog.SetLevel(beelog.LevelWarning)

�����Ļ��Ͳ������С�����level����־����־���������£�

LevelTrace��LevelDebug��LevelInfo��LevelWarning��LevelError��LevelCritical

�û����Ը��ݲ�ͬ�ļ��������ͬ�Ĵ�����Ϣ������������ʾ��


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
