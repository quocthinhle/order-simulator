import { Component, signal } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-auth-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  credentialForm: FormGroup;

  x = signal(5);

  constructor() {
    this.credentialForm = new FormGroup({
      password: new FormControl('', [Validators.required]),
      username: new FormControl('', [Validators.required, Validators.minLength(4)]),
    });
  }

  onLoginFormSubmit() {
    console.log(this.credentialForm.value);
  }

}
