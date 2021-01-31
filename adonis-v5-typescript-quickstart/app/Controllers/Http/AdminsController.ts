// import { HttpContextContract } from '@ioc:Adonis/Core/HttpContext'

import { Request } from "@adonisjs/core/build/standalone";
import View from "@ioc:Adonis/Core/View";
import { HttpContextContract } from '@ioc:Adonis/Core/HttpContext'

export default class AdminsController {
    public async index() {
        return View.render('admin')
    }
    public async connect({ request }: HttpContextContract) {
        console.log(request.all())
        return 'connect admin'
    }
}
