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
      });
  }

  getMovie(ID: string){
    var config = {
      headers:
          {
              'Content-Type': 'application/json',
              'Access-Control-Allow-Credentials': 'true'
          }
  }
    this.httpClient.get('localhost:3000/api/getMovies/',config).subscribe(
      res => {
        this.test = 'TRUE';
        this.title = res['data'];
      }
    );
      
  }
  ngOnInit() {
    
  }

}
