$(document).ready(function() {


Highcharts.setOptions({
  global : {
    useUTC : false
  }
});

$('#data').highcharts('StockChart', {
  chart : {
    events : {
        load : function() {

          var log = $("#log");
          var priceData = this.series[0]
          var volData = this.series[1]
          var conn;

          function appendLog(msg) {
              var d = log[0]
              var doScroll = d.scrollTop == d.scrollHeight - d.clientHeight;
              msg.appendTo(log)
              if (doScroll) {
                  d.scrollTop = d.scrollHeight - d.clientHeight;
              }
          }


          if (window["WebSocket"]) {
              conn = new WebSocket("ws://"+window.location.hostname+":8080/ws");
              conn.onopen = function(evt) {
                  appendLog($("<div>Connection opened.</div>"))
              }
              conn.onclose = function(evt) {
                  appendLog($("<div>Connection closed.</div>"))
              }
              conn.onmessage = function(evt) {
                  appendLog($("<div/>").text(evt.data))
                  var obj = JSON.parse(evt.data)
                  var time = (new Date()).getTime()
                  var price = [ time , obj.price];
                  var volume = [ time , obj.amount];
                  priceData.addPoint(price,true,false);
                  volData.addPoint(volume,true,false);
              }
          } else {
              appendLog($("<div><b>Your browser does not support WebSockets.</b></div>"))
          }
        }
      }
  },
  rangeSelector: {
    buttons: [{
      count: 1,
      type: 'minute',
      text: '1M'
    }, {
      count: 5,
      type: 'minute',
      text: '5M'
    }, {
      type: 'all',
      text: 'All'
    }],
    inputEnabled: false,
    selected: 0
  },
  
  title : {
    text : 'Bitstamp'
  },
  yAxis: [{ // Primary yAxis
                labels: {
                    format: '{value}$',
                    style: {
                        color: Highcharts.getOptions().colors[0]
                    }
                },
                title: {
                    text: 'Price',
                    style: {
                        color: Highcharts.getOptions().colors[0]
                    }
                },
                height: '60%'
            }, { // Secondary yAxis
                title: {
                    text: 'Volume',
                    style: {
                        color: Highcharts.getOptions().colors[1]
                    }
                },
                labels: {
                    format: '{value} BTC',
                    style: {
                        color: Highcharts.getOptions().colors[1]
                    }
                },
                top: '65%',
                height: '35%'
  }],
  exporting: {
    enabled: false
  },
  series : [{
    name : 'Price',
    type: 'line',

    data : (function() {
      // generate an array of random data
      var data = [], time = (new Date()).getTime();
      data.push([
          time ,
          640
      ]);
      return data;
    })()
  },
  {
    name : 'Volume',
    yAxis: 1,
    zIndex: -1,
    type: 'column',
    data : (function() {
      // generate an array of random data
      var data = [], time = (new Date()).getTime();
      data.push([
          time ,
          1
      ]);
      return data;
    })()
  }]
});



});