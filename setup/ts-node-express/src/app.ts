import createError from 'http-errors'
import express from 'express'
import path from 'path'
import cookieParser from 'cookie-parser'
import logger from 'morgan'

import indexRouter from './routes/index'
import usersRouter from './routes/users'

const app = express()

// view engine setup
app.set('views', path.join(__dirname, 'views'))
app.set('view engine', 'ejs')

app.use(logger('dev'))
app.use(express.json())
app.use(express.urlencoded({ extended: false }))
app.use(cookieParser())
app.use(express.static(path.join(__dirname, 'public')))

app.use('/', indexRouter)
app.use('/users', usersRouter)

// catch 404 and forward to error handler
app.use((request: express.Request, response: express.Response, next) => {
  next(createError(404))
})

// error handler
app.use((errors: any, request: express.Request, response: express.Response, next: any) => {
  // set locals, only providing error in development
  response.locals.message = errors.message
  response.locals.error = request.app.get('env') === 'development' ? errors : {}

  // render the error page
  response.status(errors.status || 500)
  response.render('error')
})

export default app
