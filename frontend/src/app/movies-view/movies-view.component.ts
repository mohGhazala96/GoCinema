import { Component, OnInit } from '@angular/core';
import {Router, ActivatedRoute } from '@angular/router';
import {HttpClient} from '@angular/common/http';

export class Movie {
  Id : Number ;
  Vote_average : Number;
  Title : string;
  Poster_path : string;
  Overview : string;
  Release_date : string
};

@Component({
  selector: 'app-movies-view',
  templateUrl: './movies-view.component.html',
  styleUrls: ['./movies-view.component.css']
})
export class MoviesViewComponent implements OnInit {

  movies : Movie[];

  constructor(private router : Router, private activatedRoute: ActivatedRoute, private httpClient: HttpClient) { 


  }

  getMovies() {

    var config = {
      headers:
          {
              'Content-Type': 'application/json'
          }
    }  

    this.httpClient.get('http://localhost:3000/api/getMovies').subscribe(

      res => {
        

        for( var i = 0; i < 20; i++) {

          this.movies[i] = new Movie() 
          this.movies[i].Id = res['Movies'][i].Id
          this.movies[i].Vote_average = res['Movies'][i].Vote_average
          this.movies[i].Title = res['Movies'][i].Title
          this.movies[i].Poster_path = res['Movies'][i].Poster_path
          this.movies[i].Overview = res['Movies'][i].Overview
          this.movies[i].Release_date = res['Movies'][i].Release_date

        }
      }
    );

  }

  reserveSeat( movieId: Number){

   window.location.href = "/reserve?movie_id=" + movieId
  }

  ngOnInit() {
    this.movies = new Array(20)
    this.getMovies()
    console.log(this.movies)
  }

}
