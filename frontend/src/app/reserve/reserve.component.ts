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

  constructor(private activatedRoute: ActivatedRoute, private httpClient: HttpClient) {
    this.activatedRoute.queryParams.subscribe(params => {
          this.movie_id = params['movie_id'];
          console.log(this.movie_id); // Print the parameter to the console. 
          this.getMovie(this.movie_id);
          this.sendPing(this.movie_id);
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
