var gulp = require('gulp');
var run = require('gulp-run');
var plumber = require('gulp-plumber');
var notify = require('gulp-notify');

gulp.task('test', function() {
    run('go test', {cwd: 'missle'}).exec()
    .on('error', notify.onError({
        title: "Crap",
        message: "Your tests failed, Jeffrey!"
    }))
    .pipe(notify({
        title: "Success",
        message: "All tests have returned green!"
    }));
});

gulp.task('watch', function() {
    gulp.watch('missle/**/*.go', ['test'])
    .on('change', function(event) {
      console.log('File ' + event.path + ' was ' + event.type + ', running tasks...');
    });;
});

// gulp.task('watch', function() {
//     gulp.watch('missle/**/*.go', function(event) {
//         console.log('File ' + event.path + ' was ' + event.type + ', running tasks...');
//     });
// });

gulp.task('default', ['test', 'watch']);