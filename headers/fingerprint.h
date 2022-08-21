#ifndef _FINGER_NEBUIA_H
#define _FINGER_NEBUIA_H

#include <stdbool.h>

#ifdef __cplusplus
extern "C"
{
#endif
    typedef struct FingerDetections
    {
        int detections[4][4];
    } FingerDetections;

    typedef struct Fingerprint
    {
        int length;
        unsigned char *image;
    } Fingerprint;

    void initFingersModel();
    FingerDetections *detectFingerprints(unsigned char *image, int length);
    bool extractFingerPrints(unsigned char *image, int length);
    int getNFiq(unsigned char *image, int length);
    int generateWSQ(unsigned char *image, int length);
#ifdef __cplusplus
}
#endif

#endif /* !_FINGER_NEBUIA_H */