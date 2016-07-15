function initcanvas(){
 
 $(document).ready(function() {
  
//var items = [];
 
$.get( "/gettitles", function( data ) {
   for(var item in data){
     
     var href = "/get/" +  data[item].title;
     var t = $("<a></a>").text(data[item].title).attr("href", href);   
     $("#links").append(t);
     }

 if(!$('#myCanvas').tagcanvas({
          textColour: '#ff0000',
          textFont: null,
          weight:true,
          textHeight: 20,
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
   //alert(data[0].title);
});
 
}

function retrievetags(){
     $.get("http://localhost:3000/retievetags",function(data,status){
     	tags = data
     }); 
     
} 