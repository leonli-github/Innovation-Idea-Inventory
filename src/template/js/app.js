var idea={};

function initcanvas(){
 
 $(document).ready(function() {
  
//var items = [];
 
$.get( "/gettitles", function( data ) {
   for(var item in data){
     
     var fontno = Math.floor(Math.random() * 30) + 10;
     var fontsize = fontno.toString() + "pt";
     var href = "/get/" +  data[item].projecttitle;
     idea[data[item].projecttitle] = data[item].description;
     //var t = $("<a></a>").text(data[item].projecttitle).css({"font-size":fontsize}).attr("href", href); 
     var t = $("<a></a>").text(data[item].projecttitle).attr("id",data[item].projecttitle).css({"font-size":fontsize}).on( "click", function(event) {
  
       event.preventDefault();
       //alert(idea[$(this).text()]);
       OpenDialog($(this).text(),idea[$(this).text()]);
  
});   
     $("#links").append(t);
     }

 if(!$('#myCanvas').tagcanvas({
          textColour: null,
          textFont: null,
          weight:true,
          weightMode:"both",
          //textHeight: 20,
          outlineColour: '#ff00ff',
          reverse: true,
          depth: 0.8,
          initial: [0.060, 0.150],
          maxSpeed: 0.05

        },'tags')) {
          // something went wrong, hide the canvas container
          $('#myCanvasContainer').hide();
        }
      });
   //alert(data[0].projecttitle);
});
 
}

function retrievetags(){
     $.get("http://localhost:3000/retievetags",function(data,status){
     	tags = data
     }); 
     
} 

function OpenDialog(name,desc){
     
     $("#dialog").text(desc);
     $("#dialog").dialog('option','title',name);
     $( "#dialog" ).dialog( "open" );
}