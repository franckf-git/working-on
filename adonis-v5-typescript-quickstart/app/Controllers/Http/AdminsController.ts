import { Request } from "@adonisjs/core/build/standalone";
import View from "@ioc:Adonis/Core/View";
import { HttpContextContract } from '@ioc:Adonis/Core/HttpContext'
import { schema, validator } from '@ioc:Adonis/Core/Validator'

export default class AdminsController {
    public async index() {
        return View.render('admin')
    }
    public async connect({ request }: HttpContextContract) {
        const postSchema = schema.create({ username: schema.string(), password: schema.string() })
        const data = await request.validate({ schema: postSchema, cacheKey: request.url() })

        console.log(data)
        return 'connect admin'
    }
}
