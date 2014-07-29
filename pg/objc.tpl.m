//
//  {{.FileNameBase}}.m
//  {{.ClassName}}
//
//  Created by generator on 14/6/29.
//  Copyright (c) 2014 Wan Wei. All rights reserved.
//

@interface Car()
@property BOOL isSelected;
@property (nonatomic)  int timeOut;
@end

@implementation Car {
    SKLabelNode* _countDown;
    NSTimer* _countDownTimer;
}

+(instancetype) carWithId:(int64_t) carId IsLeft:(BOOL)isLeft {
    CGSize size = CGSizeMake(15, 10);
    Car *car = [Car spriteNodeWithColor:isLeft?[UIColor redColor]:[UIColor greenColor] size:size];
    car.physicsBody = [SKPhysicsBody bodyWithRectangleOfSize:size];
    car.physicsBody.friction = 1.0f;
    car.physicsBody.usesPreciseCollisionDetection = YES;
    car.isSelected = NO;
    car.isLeft = isLeft;
    [car setUserInteractionEnabled:YES];
    return car;
}

-(instancetype)initWithColor:(UIColor *)color size:(CGSize)size {
    if (self = [super initWithColor:color size:size]) {
        _timeOut = 5;
        _countDown = [SKLabelNode labelNodeWithFontNamed:@"System"];
        _countDown.position = CGPointMake(0, 30);
        _countDown.hidden = YES;
        [self addChild:_countDown];
    }
    
    return self;
}
