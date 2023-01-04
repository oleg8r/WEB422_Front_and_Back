import { Component } from '@angular/core';
import {Drug} from "../models/drug.model";
import axios from "axios";

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})


export class AppComponent {
  title = 'drugs-app';
  newDrug: Drug = {};
  drugs: Drug[] = [];

 //  drugs: Drug[] = [
 //    {
 //      uuid:'0001',
 //      name: "Drug1",
 //      summary: 'summary',
 //      dateAdded:'Jan 1, 2022',
 //      quantity: 0;
 //    }
 // ]
  ngOnInit(){
    axios.get("http://localhost:3000/read")
      .then((response)=>{
        console.log(typeof(response.data))
        this.drugs = <Drug[]>JSON.parse(JSON.stringify(response.data))
        console.log(this.drugs)
      })
  }

  addDrug(): void {
    console.log(this.newDrug)
    axios.post('http://localhost:3000/create', {
      name: this.newDrug.name,
      description: this.newDrug.description,
      quantity: this.newDrug.quantity
    })
      .then((response) =>{
        this.drugs.push(<Drug>JSON.parse(JSON.stringify(response.data)))
        console.log(response);
      })
      .catch(function (error) {
        console.log(error);
      });
  }

  updateDrug(id?: number, quantity?: number): void {
    console.log(this.newDrug)

    let formData = new FormData()
    formData.append('id', id?.toString() || '')
    formData.append('quantity', quantity?.toString() || '')

    axios.post('http://localhost:3000/update', formData)

      .then((response) => {
        //this.drugs.push(<Drug>JSON.parse(JSON.stringify(response.data)))
        console.log(response);
      })

      .catch(function (error) {
        console.log(error);
      });
  }

  deleteDrug(i:number) {
    console.log(i)
    axios.delete('http://localhost:3000/delete?id='+ this.drugs[i].ID)
      .then((response) =>{
        console.log(response);
        this.drugs.splice(i, 1)
      })

      .catch(function (error) {
        console.log(error);
      })


  }

}
