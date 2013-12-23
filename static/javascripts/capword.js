'use strict';

var capwordApp = angular.module('capwordApp', ['capwordService']);
var capwordService = angular.module('capwordService', ['ngResource']);

capwordService.factory("WordsManager", function($resource) {
  var Words = $resource('/words/index', {
    format: 'json'
  });

  var WordsManager = {
    current: [],
    get: function() {
      var self = this;
      Words.get({}, function(words, header) {
        angular.forEach(words.Elements, function(word, _) {
          self.current.push(word);
        });
      });
      return this.current;
    }
  };

  return WordsManager;
});

capwordApp.controller('capwordController', function($scope, WordsManager) {
  $scope.words = WordsManager.get();
});