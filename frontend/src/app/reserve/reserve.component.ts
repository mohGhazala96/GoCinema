import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import {HttpClient} from '@angular/common/http';
import { ThrowStmt } from '@angular/compiler';

@Component({
  selector: 'app-reserve',
  templateUrl: './reserve.component.html',
  styleUrls: ['./reserve.component.css']
})
export class ReserveComponent implements OnInit {
  movie_id: string;
  title: string;
  test: string;
  rowsLength = 10;
  columnsLength = 10;
  seatsLayout = [];
  reservedSeats = []
  rowLetters = ['A','B','C','D','E','F','G','H','I','J']
  availableTime = ['1:00pm - 3:00pm','4:00pm - 6:00pm','7:00pm - 9:00pm','10:00pm - 12:00am']
  availableTimeToSend = ['1','4','7','10']
  selectedTime = this.availableTime[0]
  selectedDate = ""
  error=""
  enteredEmail = ""

  currentMovie = {
    Title: "",
    Overview: "",
    Poster_path: "",
    Release_date: "",
    Vote_average: 0.0,
    hall: 0
  }

  showSeats = false;


  newReservation = []

  constructor(private activatedRoute: ActivatedRoute, private httpClient: HttpClient) {
    console.log("INIT ")
    for(var i=0;i<this.rowsLength;i++){
      this.seatsLayout[i]= [];
      for(var j=0;j<this.columnsLength;j++){
        this.seatsLayout[i][j] = 0;
      }
    }
    for(var i=0;i<this.rowsLength;i++){
      this.newReservation[i]= [];
      for(var j=0;j<this.columnsLength;j++){
        this.newReservation[i][j] = 0;
      }
    }
    this.activatedRoute.queryParams.subscribe(params => {
          this.movie_id = params['movie_id'];
          console.log(this.movie_id); // Print the parameter to the console. 
          this.getMovie(this.movie_id);
          // this.getReservations(this.movie_id,"1","2018-12-5")

      });
  }

  selectTime(event){
    this.selectedTime = event
    console.log(this.selectedTime)
  }

  selectDate(event){
    this.selectedDate = event
    console.log(this.selectedDate)
  }

  pingSeats(){
    console.log(this.selectedTime)
    console.log(this.selectedDate)

    if(this.selectedTime == "" || this.selectedDate== ""){
      this.error= "Please select date/time"
    }else{
      this.getReservations(this.movie_id,this.selectedTime,this.selectedDate)
    }
  }

  getMovie(ID: string){
    var config = {
      headers:
          {
              'Content-Type': 'application/json',
              // 'Access-Control-Allow-Credentials': 'true'
          }
  }
    this.httpClient.get('http://localhost:3000/api/getMovie?movie_id='+ID,config).subscribe(
      res => {
        console.log(res)
        var movie = res['Movies'][0];
        this.currentMovie.Title = movie.Title;
        this.currentMovie.Overview = movie.Overview;
        this.currentMovie.Release_date = movie.Release_date;
        this.currentMovie.Poster_path = movie.Poster_path;
        this.currentMovie.Vote_average = movie.Vote_average;
        this.currentMovie.hall = movie.Hall_Id

      }
    );
      
  }

  

  getReservations(ID: string, timing: string, day: string){
    var timeToCheck = this.availableTimeToSend[this.availableTime.indexOf(timing)]
    console.log(ID+" "+timeToCheck+" "+day)
    var config = {
      headers:
          {
              'Content-Type': 'application/json',
              // 'Access-Control-Allow-Credentials': 'true'
          }
  }
    this.httpClient.get('http://localhost:3000/api/checkSeats?movieId='+ID+"&timing="+timeToCheck+"&day="+day,config).subscribe(
      res => {
        if(res!=null){
        console.log(res)
        var resLength = Object.keys(res).length
          for(var i=0;i<resLength;i++){
            var seat = res[i].toString();
            var outer = this.rowLetters.indexOf(seat[0])
            var inner = parseInt(seat[1])
            this.seatsLayout[outer][inner]=1
          }
        }
        this.showSeats =true;
      }
    );
  }

  enterEmail(event){
    this.enteredEmail = event;
    console.log(this.enteredEmail)
  }

  pressButton(outer: number, inner: number){
    if(this.reservedSeats.includes(this.rowLetters[outer]+""+inner)){
      var indexOfItem = this.reservedSeats.indexOf(this.rowLetters[outer]+""+inner);
      this.newReservation[outer][inner]=0;
      console.log(indexOfItem)
      this.reservedSeats.splice(indexOfItem,1);
    }else{
    console.log("Outer is "+this.rowLetters[outer] +"inner is "+inner);
    this.reservedSeats.push(this.rowLetters[outer]+""+inner);
    this.newReservation[outer][inner]=1;
    console.log(this.reservedSeats)
    }
  }

  reserveSeats(){
    if(this.enteredEmail == "" || this.selectedDate == "" || this.selectedTime=="" || this.reservedSeats.length ==0 ){
      this.error = "Please make sure you entered all inputs and selected seats"
    }else{
      this.error = ""
      var timeToCheck = this.availableTimeToSend[this.availableTime.indexOf(this.selectedTime)]
      console.log(this.enteredEmail + " " +this.selectedDate +" "+this.selectedTime + " "+ this.reservedSeats)
      var config = {
        headers:
            {
                'Content-Type': 'application/json',
                // 'Access-Control-Allow-Credentials': 'true'
            }
    }
  
  
      var newReservation = JSON.stringify
        ({
            // userid:localStorage.getItem("user"),
            "Hall": this.currentMovie.hall,
            "Movie": parseInt(this.movie_id),
            "Useremail":this.enteredEmail,
            "Day":this.selectedDate,
            "Timing": parseInt(timeToCheck),
            "Seats":this.reservedSeats
        });  
        this.httpClient.post('http://localhost:3000/api/insertReservation',newReservation,config).subscribe(
      res => {
        console.log(newReservation)
        console.log("Reservation success")
        console.log(res)
      }
    );    

    }
  }

  sendPing(ID: string){
    var config = {
      headers:
          {
              'Content-Type': 'application/json',
              // 'Access-Control-Allow-Credentials': 'true'
          }
  }


  var newReservation = JSON.stringify
            ({
                // userid:localStorage.getItem("user"),
                title: "ping successful"
            });
    this.httpClient.post('http://localhost:3000/api/addReservation/',newReservation,config).subscribe(
      res => {
        this.test = 'TRUE';
      }
    );
      
  }


  ngOnInit() {

  }

}
