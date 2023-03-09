/**
 *formatTime参数   时间戳转换为时间格式
 * 第一个参数：时间戳，（必填项）
 * 第二个参数：日期间隔符，默认为 -，(选填)
 * 第三个参数：时分秒间隔符，默认为 :，(选填)
 * 第四个参数：是否携带时分秒，默认为 true，(选填)
 */
// let time = formatTime(1579258960);
// console.log(time);//2020-01-17 19:02:40

function formatTime(timestamp, separator = '-', timeS = ':', flag = true) {
	let str = ''
	let date = new Date(timestamp * 1000) //时间戳为10位需*1000，时间戳为13位的话不需乘1000
	let Y = date.getFullYear() + separator
	let M = autoChange(date.getMonth() + 1) + separator //计算机的月份是从0开始滴，需要+1
	let D = autoChange(date.getDate())
	str = Y + M + D
	if (flag) {
		let h = autoChange(date.getHours()) + timeS
		let m = autoChange(date.getMinutes()) + timeS
		let s = autoChange(date.getSeconds())
		let timeStr = h + m + s
		str += ' '
		str += timeStr
	}
	return str
}
function autoChange(num) {
	if (num < 10) {
		return '0' + num
	} else {
		return num
	}
}

/*
 ** 时间戳显示为多少分钟前，多少天前的处理
 ** eg.
 ** console.log(dateDiff(1411111111111));  // 2014年09月19日
 ** console.log(dateDiff(1481111111111));  // 9月前
 ** console.log(dateDiff(1499911111111));  // 2月前
 ** console.log(dateDiff(1503211111111));  // 3周前
 ** console.log(dateDiff(1505283100802));  // 1分钟前
 */
function dateDiff(timestamp) {
	// 补全为13位
	var arrTimestamp = (timestamp + '').split('')
	for (var start = 0; start < 13; start++) {
		if (!arrTimestamp[start]) {
			arrTimestamp[start] = '0'
		}
	}
	timestamp = arrTimestamp.join('') * 1

	var minute = 1000 * 60
	var hour = minute * 60
	var day = hour * 24
	// var halfamonth = day * 15
	var month = day * 30
	var now = new Date().getTime()
	var diffValue = now - timestamp

	// 如果本地时间反而小于变量时间
	if (diffValue < 0) {
		return '不久前'
	}

	// 计算差异时间的量级
	var monthC = diffValue / month
	var weekC = diffValue / (7 * day)
	var dayC = diffValue / day
	var hourC = diffValue / hour
	var minC = diffValue / minute

	// 数值补0方法
	var zero = function (value) {
		if (value < 10) {
			return '0' + value
		}
		return value
	}

	// 使用
	if (monthC > 12) {
		// 超过1年，直接显示年月日
		return (function () {
			var date = new Date(timestamp)
			return (
				date.getFullYear() +
				'年' +
				zero(date.getMonth() + 1) +
				'月' +
				zero(date.getDate()) +
				'日'
			)
		})()
	} else if (monthC >= 1) {
		return parseInt(monthC) + '月前'
	} else if (weekC >= 1) {
		return parseInt(weekC) + '周前'
	} else if (dayC >= 1) {
		return parseInt(dayC) + '天前'
	} else if (hourC >= 1) {
		return parseInt(hourC) + '小时前'
	} else if (minC >= 1) {
		return parseInt(minC) + '分钟前'
	}
	return '刚刚'
}

/**
	 * 获取系统时间时间戳
	 * @param {*} date 时间
	 */
function getTimeStamp() {
	return new Date().getTime().toString().substring(0, 10)
}




function remainTime(timeStamp) {
	let remain = new Date().getTime() - timeStamp * 1000
	let day = parseInt(remain / 1000 / 60 / 60 / 24) + 1
	return day + "天前";
}
function deadline(timeStamp) {
	let remain = timeStamp * 1000 - new Date().getTime()
	if (remain <= 0) {
		return "超时"
	}
	else {
		let days = parseInt(remain / 1000 / 60 / 60 / 24) + 1
		return "还剩" + days + "天";
	}

}
function dataTimeFormatter(v){
	let date = new Date(v);
	let y = date.getFullYear();
	let MM = date.getMonth() + 1;
	MM = MM < 10 ? "0" + MM : MM;
	let d = date.getDate();
	d = d < 10 ? "0" + d : d;
	let h = date.getHours();
	h = h < 10 ? "0" + h : h;
	let m = date.getMinutes();
	m = m < 10 ? "0" + m : m;
	let s = date.getSeconds();
	s = s < 10 ? "0" + s : s;
	return y + "-" + MM + "-" + d + " " + h + ":" + m + ":" + s;
  }
export { formatTime, dateDiff, getTimeStamp, remainTime, deadline,dataTimeFormatter }