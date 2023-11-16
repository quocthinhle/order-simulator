import { NgModule } from "@angular/core";
import { RouterModule, Route } from "@angular/router";

const ROUTES: Route[] = [
    {
        path: 'auth',
        loadChildren: () => import('./internal/screens/auth/auth-module').then(m => m.AuthModule),
    },
    {

    },
];

@NgModule({
    imports: [
        RouterModule.forRoot(ROUTES),
    ]
})
export class AppRoutingModule {}