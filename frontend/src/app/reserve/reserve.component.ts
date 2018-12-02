import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import {HttpClient} from '@angular/common/http';

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
          // this.getMovie(this.movie_id);
          // this.sendPing(this.movie_id);
          // this.testInsert(this.movie_id);
      });
  }

  getMovie(ID: string){
    var config = {
      headers:
          {
              'Content-Type': 'application/json',
              // 'Access-Control-Allow-Credentials': 'true'
          }
  }
    this.httpClient.get('http://localhost:3000/api/getMovies/',config).subscribe(
      res => {
        this.title = res['Movies'][0]['Title'];
        console.log(res['Movies'])
      }
    );
      
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

  testInsert(ID: string){
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
                Hall: 1,
                Seats: ['A1','A2'],
                Movie: 1,
                Useremail: "hello@live.com",
                Timing: 2
            });
    this.httpClient.post('http://localhost:3000/api/insertReservation',newReservation,config).subscribe(
      res => {
        this.test = 'TRUE';
      }
    );
      
  }

  ngOnInit() {

  }

}
