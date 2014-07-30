<?php

/**
* Engine
*type Player struct {
*    UserId int64
*    Name   string
*    Age    int
*    Points int
*}
*
*type PlayerService interface {
*    GetPlayer(userId int64) (player *Player)
*    GetOnlinePlayers() (players []*Player)
*}
*/

$messages = array(
    'Player' => array(
        'UserId' => 'int64',
        'Name' => 'string',
        'Age' => 'int',
        'Points' => 'int',
    ),
);

$services = array(
    'GetPlayer' => array(
        'In' => array(
            'userId' => '[]int64',
        ),
        'Out' => array(),
    ),
);

class Engine
{
    private $objcTypes = array('string'=>'NSString*');

    public function Make($value='')
    {
        # code...
    }
}