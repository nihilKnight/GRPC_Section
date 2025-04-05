PROGRAM MotorControl
VAR
    SpeedSetpoint: INT := 1000;
    ActualSpeed: INT;
    ErrorCode: WORD;
END_VAR

IF SpeedSetpoint > 0 THEN
    StartMotor(SpeedSetpoint);
    ActualSpeed := GetMotorSpeed();
    
    IF ActualSpeed < SpeedSetpoint * 0.9 THEN
        ErrorCode := 16#1001;
    END_IF
END_IF
END_PROGRAM